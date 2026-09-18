package scraper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

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

func denarioTime(t *testing.T, value string) time.Time {
	t.Helper()
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		t.Fatalf("parse %s: %v", value, err)
	}
	return parsed
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

func TestDenarioUpdateValuesSendsPricesAndReserves(t *testing.T) {
	server := fakeDenarioAPI(t, map[string]string{
		"/price/goldcoin/latest/CHF":   `{"priceAsk":"3796.19812","priceBid":"3644.35019","currency":"CHF","priceDate":"2026-09-18T04:34:02+02:00"}`,
		"/price/silvercoin/latest/USD": `{"priceAsk":"74.3732","priceBid":"71.1","currency":"USD","priceDate":"2026-09-18T04:34:05+02:00"}`,
		"/reserve/goldcoins":           `{"date":"2026-09-18T04:34:29+02:00","amountOunces":"46.30234"}`,
	})
	scraper := testDenarioScraper(server, testDenarioAuthKey,
		[]string{"goldcoin/latest/CHF", "silvercoin/latest/USD"},
		[]string{"goldcoins"},
	)

	if err := scraper.UpdateValues(); err != nil {
		t.Fatalf("UpdateValues: %v", err)
	}

	want := []DenarioQuote{
		{Key: "goldcoin/latest/CHF", Value: 3796.19812, Time: denarioTime(t, "2026-09-18T04:34:02+02:00"), Type: DenarioPrice},
		{Key: "silvercoin/latest/USD", Value: 74.3732, Time: denarioTime(t, "2026-09-18T04:34:05+02:00"), Type: DenarioPrice},
		{Key: "goldcoins", Value: 46.30234, Time: denarioTime(t, "2026-09-18T04:34:29+02:00"), Type: DenarioReserve},
	}
	got := sentDenarioQuotes(t, scraper)
	if len(got) != len(want) {
		t.Errorf("sent %d quotes, want %d: %+v", len(got), len(want), got)
	}
	for _, quote := range want {
		if !got[quote.Key].Time.Equal(quote.Time) {
			t.Errorf("%s time = %v, want %v", quote.Key, got[quote.Key].Time, quote.Time)
		}
		if got[quote.Key].Value != quote.Value || got[quote.Key].Type != quote.Type {
			t.Errorf("sent %+v, want %+v", got[quote.Key], quote)
		}
	}
}

// The API reports when it priced the asset, so a backend that stopped updating
// can be kept off the oracle.
func TestDenarioSkipsStaleValues(t *testing.T) {
	fresh := time.Now().Add(-time.Minute)
	stale := time.Now().Add(-2 * time.Hour)
	server := fakeDenarioAPI(t, map[string]string{
		"/price/goldcoin/latest/CHF": `{"priceAsk":"3796.19812","priceDate":"` + fresh.Format(time.RFC3339) + `"}`,
		"/price/goldcoin/latest/USD": `{"priceAsk":"4612.03927","priceDate":"` + stale.Format(time.RFC3339) + `"}`,
	})
	scraper := testDenarioScraper(server, testDenarioAuthKey,
		[]string{"goldcoin/latest/CHF", "goldcoin/latest/USD"}, nil)
	scraper.maxValueAge = 30 * time.Minute

	if err := scraper.UpdateValues(); err != nil {
		t.Fatalf("UpdateValues: %v", err)
	}

	got := sentDenarioQuotes(t, scraper)
	if _, sent := got["goldcoin/latest/USD"]; sent {
		t.Error("a 2h old price was sent although maxValueAge is 30m")
	}
	if _, sent := got["goldcoin/latest/CHF"]; !sent {
		t.Error("the fresh price was not sent")
	}
}

// Age limit off (the default): old values still go to the oracle.
func TestDenarioKeepsOldValuesWithoutAgeLimit(t *testing.T) {
	server := fakeDenarioAPI(t, map[string]string{
		"/price/goldcoin/latest/CHF": `{"priceAsk":"3796.19812","priceDate":"2020-01-01T00:00:00+02:00"}`,
		// No timestamp at all, e.g. if the API changes shape.
		"/reserve/goldcoins": `{"amountOunces":"46.30234"}`,
	})
	scraper := testDenarioScraper(server, testDenarioAuthKey,
		[]string{"goldcoin/latest/CHF"}, []string{"goldcoins"})

	if err := scraper.UpdateValues(); err != nil {
		t.Fatalf("UpdateValues: %v", err)
	}

	got := sentDenarioQuotes(t, scraper)
	if len(got) != 2 {
		t.Errorf("sent %d quotes, want 2: %+v", len(got), got)
	}
	if !got["goldcoins"].Time.IsZero() {
		t.Errorf("missing timestamp became %v, want zero", got["goldcoins"].Time)
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

	// The price route is <coin>/latest/<currency>, a short path answers 404.
	for _, asset := range c.Prices {
		if strings.Count(asset, "/") != 2 || !strings.Contains(asset, "/latest/") {
			t.Errorf("price asset %q is not a <coin>/latest/<currency> path", asset)
		}
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
		value, valueTime, err := scraper.getValue(DENARIO_PRICE_URL+"/"+asset, DENARIO_PRICE_FIELD, DENARIO_PRICE_TIME_FIELD)
		if err != nil {
			t.Errorf("price %s: %v", asset, err)
			continue
		}
		t.Logf("price %s = %v (%s, %v old)", asset, value, valueTime.Format(time.RFC3339), time.Since(valueTime).Truncate(time.Second))
	}
	for _, asset := range c.Reserves {
		value, valueTime, err := scraper.getValue(DENARIO_RESERVE_URL+"/"+asset, DENARIO_RESERVE_FIELD, DENARIO_RESERVE_TIME_FIELD)
		if err != nil {
			t.Errorf("reserve %s: %v", asset, err)
			continue
		}
		t.Logf("reserve %s = %v (%s, %v old)", asset, value, valueTime.Format(time.RFC3339), time.Since(valueTime).Truncate(time.Second))
	}
}
