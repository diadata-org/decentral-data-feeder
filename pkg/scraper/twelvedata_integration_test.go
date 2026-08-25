package scraper

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/diadata-org/decentral-data-feeder/pkg/models"
	utils "github.com/diadata-org/decentral-data-feeder/pkg/utils"
)

// Tickers that name a commodity on twelvedata but also match an unrelated
// equity, so they resolve to the wrong instrument unless the request pins the
// exchange. HG1 is Copper Spot but also Homag Group AG on Frankfurt, XAU is
// Gold Spot but also a gold/silver equity index.
var commodityTickersMisreadAsStocks = []string{"HG1", "XAU"}

// Guards the config itself, so it runs without an API key. The commodities loop
// is the only one that sends exchange=commodity, so these tickers have to live
// under Commodities to resolve correctly.
func TestCommodityTickersAreNotConfiguredAsStocks(t *testing.T) {
	raw, err := os.ReadFile("../../config/rwa/rwaConfig.json")
	if err != nil {
		t.Fatalf("read rwaConfig.json: %v", err)
	}

	var c models.RWAConfig
	if err := json.Unmarshal(raw, &c); err != nil {
		t.Fatalf("unmarshal rwaConfig.json: %v", err)
	}

	for _, ticker := range commodityTickersMisreadAsStocks {
		if contains(c.Stocks, ticker) {
			t.Errorf("%s is configured under Stocks, it resolves to an unrelated equity there", ticker)
		}
	}

	if !contains(c.Commodities, "HG1") {
		t.Error("HG1 missing from Commodities")
	}
	// Gold is carried as the pair, not the bare ticker.
	if !contains(c.Commodities, "XAU/USD") {
		t.Error("XAU/USD missing from Commodities")
	}
}

// Exercises the real update path: the diadata API first, twelvedata as
// fallback. Requires a live key:
//
//	TWELVEDATA_API_KEY=... go test ./pkg/scraper/ -run TestRwaSymbolResolves -v -count=1
func TestRwaSymbolResolves(t *testing.T) {
	apiKey := os.Getenv("TWELVEDATA_API_KEY")
	if apiKey == "" {
		t.Skip("TWELVEDATA_API_KEY not set")
	}

	scraper := &TwelvedataScraper{apiKey: apiKey}

	// The configured path. diadata has no HG1 under Commodities, so this always
	// falls back to twelvedata, where exchange=commodity pins it to copper.
	t.Run("HG1_commodity", func(t *testing.T) {
		quote, err := getRwaQuoteFromDia("HG1", "Commodities")
		t.Logf("dia  rwa/Commodities/HG1 : price=%v err=%v", quote.Price, err)

		commodity, err := scraper.getTwelveQuote("HG1", true)
		if err != nil {
			t.Fatalf("getTwelveQuote: %v", err)
		}
		t.Logf("td   quote+commodity     : name=%q close=%q", commodity.Name, commodity.Price)

		if commodity.Name != "Copper Spot" {
			t.Errorf("HG1 resolved to %q, want Copper Spot", commodity.Name)
		}
		if commodity.Price == "" || commodity.Price == "0" {
			t.Errorf("no usable price for HG1: %q", commodity.Price)
		}
	})

	// Documents why HG1 must not sit under Stocks: that loop calls
	// price?symbol=HG1 with no exchange, which returns a German equity in EUR.
	t.Run("HG1_as_stock_resolves_to_the_wrong_instrument", func(t *testing.T) {
		body, _, err := utils.GetRequest(twelvedataApiBaseString + "quote?symbol=HG1&apikey=" + apiKey)
		if err != nil {
			t.Fatalf("quote?symbol=HG1: %v", err)
		}
		t.Logf("td   quote?symbol=HG1    : %s", truncate(string(body), 200))

		if strings.Contains(string(body), `"name":"Copper Spot"`) {
			t.Log("twelvedata now defaults the bare HG1 ticker to copper, the Stocks/Commodities split may no longer matter")
		}
	})

	t.Run("CPER_etf", func(t *testing.T) {
		quote, err := getRwaQuoteFromDia("CPER", "ETF")
		t.Logf("dia  rwa/ETF/CPER        : price=%v name=%q err=%v", quote.Price, quote.Name, err)
		if err != nil || quote.Price <= 0 {
			t.Errorf("CPER does not resolve on the diadata API: price=%v err=%v", quote.Price, err)
		}

		etf, err := scraper.getTwelveQuote("CPER", false)
		if err != nil {
			t.Fatalf("getTwelveQuote: %v", err)
		}
		t.Logf("td   quote?symbol=CPER   : name=%q close=%q", etf.Name, etf.Price)

		// CPER also lists on LSE in GBp and on TSXV as a different company, so
		// confirm the unqualified lookup still lands on the NYSE ARCA fund.
		body, _, err := utils.GetRequest(twelvedataApiBaseString + "quote?symbol=CPER&apikey=" + apiKey)
		if err != nil {
			t.Fatalf("raw quote for CPER: %v", err)
		}
		if !strings.Contains(string(body), `"currency":"USD"`) {
			t.Errorf("CPER did not resolve to the USD listing: %s", truncate(string(body), 200))
		}
	})

	t.Run("XAU_control", func(t *testing.T) {
		asStock, errStock := getRwaQuoteFromDia("XAU", "Equities")
		asPair, errPair := getRwaQuoteFromDia("XAU/USD", "Commodities")
		t.Logf("dia  Equities/XAU        : price=%v err=%v", asStock.Price, errStock)
		t.Logf("dia  Commodities/XAU-USD : price=%v name=%q err=%v", asPair.Price, asPair.Name, errPair)

		if errPair != nil || asPair.Price <= 0 {
			t.Fatalf("XAU/USD does not resolve: price=%v err=%v", asPair.Price, errPair)
		}
		if errStock == nil && asStock.Price > 0 && asStock.Price > asPair.Price/10 {
			t.Errorf("Equities/XAU (%v) is close to gold (%v), the two paths may have converged",
				asStock.Price, asPair.Price)
		}
	})
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
