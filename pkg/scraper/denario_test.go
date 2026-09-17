package scraper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/diadata-org/decentral-data-feeder/pkg/models"
)

const testDenarioAuthKey = "test-token"

// Serves @responses by request path and, like the real API, answers 401 to
// requests without the token.
func fakeDenarioAPI(t *testing.T, responses map[string]string) *httptest.Server {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-AUTH-TOKEN") != testDenarioAuthKey {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"errorMessage":"No API token provided"}`))
			return
		}
		body, ok := responses[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)
	return server
}

// Builds a scraper against @server without starting its loops. The data
// channel is buffered, so UpdateValues does not block without an updater.
func testDenarioScraper(server *httptest.Server, authKey string, prices, reserves []string) *DenarioScraper {
	return &DenarioScraper{
		priceURL:    server.URL + "/price",
		reserveURL:  server.URL + "/reserve",
		authKey:     authKey,
		httpClient:  server.Client(),
		prices:      prices,
		reserves:    reserves,
		dataChannel: make(chan []byte, len(prices)+len(reserves)),
	}
}

func sentDenarioQuotes(t *testing.T, scraper *DenarioScraper) map[string]DenarioQuote {
	t.Helper()
	quotes := make(map[string]DenarioQuote)
	for len(scraper.dataChannel) > 0 {
		var quote DenarioQuote
		if err := json.Unmarshal(<-scraper.dataChannel, &quote); err != nil {
			t.Fatalf("unmarshal quote: %v", err)
		}
		quotes[quote.Key] = quote
	}
	return quotes
}

func readDenarioConfigFile(t *testing.T) models.DenarioConfig {
	t.Helper()
	raw, err := os.ReadFile("../../config/rwa/denario.json")
	if err != nil {
		t.Fatalf("read denario.json: %v", err)
	}

	// A misspelled field would silently leave the scraper without assets.
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	var c models.DenarioConfig
	if err := decoder.Decode(&c); err != nil {
		t.Fatalf("decode denario.json: %v", err)
	}
	return c
}

// Prices are the ask price and reserves the amount in ounces, the values the
// Denario oracles on Nexus publish. The asset path is the oracle key.
func TestDenarioUpdateValuesSendsPricesAndReserves(t *testing.T) {
	server := fakeDenarioAPI(t, map[string]string{
		"/price/goldcoin/latest/CHF":   `{"priceAsk": 3120.55, "priceBid": 3010.1}`,
		"/price/silvercoin/latest/USD": `{"priceAsk": "41.2", "priceBid": "38.9"}`,
		"/reserve/goldcoins":           `{"amountOunces": 1520.25}`,
	})
	scraper := testDenarioScraper(server, testDenarioAuthKey,
		[]string{"goldcoin/latest/CHF", "silvercoin/latest/USD"},
		[]string{"goldcoins"},
	)

	if err := scraper.UpdateValues(); err != nil {
		t.Fatalf("UpdateValues: %v", err)
	}

	want := []DenarioQuote{
		{Key: "goldcoin/latest/CHF", Value: 3120.55, Type: DenarioPrice},
		{Key: "silvercoin/latest/USD", Value: 41.2, Type: DenarioPrice},
		{Key: "goldcoins", Value: 1520.25, Type: DenarioReserve},
	}
	got := sentDenarioQuotes(t, scraper)
	if len(got) != len(want) {
		t.Errorf("sent %d quotes, want %d: %+v", len(got), len(want), got)
	}
	for _, quote := range want {
		if got[quote.Key] != quote {
			t.Errorf("sent %+v, want %+v", got[quote.Key], quote)
		}
	}
}

// One broken asset must not hold back the others.
func TestDenarioUpdateValuesSkipsFailingAssets(t *testing.T) {
	server := fakeDenarioAPI(t, map[string]string{
		"/price/goldcoin/latest/CHF": `{"priceAsk": 3120.55}`,
		"/price/goldcoin/latest/USD": `{"priceAsk": 0}`,
		"/price/goldcoin/latest/EUR": `{"priceBid": 2900.5}`,
		"/reserve/silvercoins":       `<html>Bad Gateway</html>`,
	})
	scraper := testDenarioScraper(server, testDenarioAuthKey,
		// silvercoin/latest/CHF is not served and answers 404.
		[]string{"goldcoin/latest/CHF", "goldcoin/latest/USD", "goldcoin/latest/EUR", "silvercoin/latest/CHF"},
		[]string{"silvercoins"},
	)

	if err := scraper.UpdateValues(); err != nil {
		t.Fatalf("UpdateValues: %v", err)
	}

	got := sentDenarioQuotes(t, scraper)
	if len(got) != 1 || got["goldcoin/latest/CHF"].Value != 3120.55 {
		t.Errorf("sent %+v, want only goldcoin/latest/CHF at 3120.55", got)
	}
}

// Without any value there is nothing to write, so the scraper must not
// trigger an oracle update.
func TestDenarioUpdateValuesErrorsWithoutValues(t *testing.T) {
	server := fakeDenarioAPI(t, map[string]string{
		"/price/goldcoin/latest/CHF": `{"priceAsk": 3120.55}`,
		"/reserve/goldcoins":         `{"amountOunces": 1520.25}`,
	})
	scraper := testDenarioScraper(server, "wrong-token", []string{"goldcoin/latest/CHF"}, []string{"goldcoins"})

	if err := scraper.UpdateValues(); err == nil {
		t.Error("UpdateValues returned no error although every request was rejected")
	}
	if n := len(scraper.dataChannel); n != 0 {
		t.Errorf("sent %d quotes, want none", n)
	}
}

func TestDenarioConfigFile(t *testing.T) {
	c := readDenarioConfigFile(t)
	if len(c.Prices) == 0 {
		t.Error("no Prices configured")
	}
	if len(c.Reserves) == 0 {
		t.Error("no Reserves configured")
	}
}

// Checks the configured assets against the real API. Requires a Denario API token:
//
//	DENARIO_AUTH_KEY=... go test ./pkg/scraper/ -run TestDenarioLiveAPI -v -count=1
func TestDenarioLiveAPI(t *testing.T) {
	authKey := os.Getenv("DENARIO_AUTH_KEY")
	if authKey == "" {
		t.Skip("DENARIO_AUTH_KEY not set")
	}

	c := readDenarioConfigFile(t)
	scraper := &DenarioScraper{
		authKey:    authKey,
		httpClient: &http.Client{Timeout: denarioRequestTimeout},
	}

	for _, asset := range c.Prices {
		value, err := scraper.getValue(DENARIO_PRICE_URL+"/"+asset, DENARIO_PRICE_FIELD)
		if err != nil {
			t.Errorf("price %s: %v", asset, err)
			continue
		}
		t.Logf("price %s = %v", asset, value)
	}
	for _, asset := range c.Reserves {
		value, err := scraper.getValue(DENARIO_RESERVE_URL+"/"+asset, DENARIO_RESERVE_FIELD)
		if err != nil {
			t.Errorf("reserve %s: %v", asset, err)
			continue
		}
		t.Logf("reserve %s = %v", asset, value)
	}
}
