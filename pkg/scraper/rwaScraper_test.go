package scraper

import (
	"encoding/json"
	"testing"
	"time"
)

// Builds a scraper that collects quotes without publishing them: the cooldown
// has not elapsed, so a pending quote is left in the batch for inspection
// instead of being flushed to the oracle.
func testCollectingScraper() *RWAWSScraper {
	return &RWAWSScraper{
		pendingQuotes:       make(map[string]RWAWSQuote),
		lastPublishedPrices: make(map[string]float64),
		lastPublishedTimes:  make(map[string]time.Time),
		deviationThresholds: make(map[string]float64),
		publishCooldown:     time.Hour,
		lastPublishTime:     time.Now(),
		commodities:         []string{"XAU/USD"},
	}
}

// Twelvedata pushes price 0 over the websocket for stale or between-session
// instruments. Publishing one writes a 0 to the oracle, which is what the XAU
// backtest surfaced.
func TestHandlePriceMessageDropsNonPositivePrice(t *testing.T) {
	for _, raw := range []string{`0`, `"0"`, `0.0`, `-1`, `"-0.5"`} {
		scraper := testCollectingScraper()

		msg := rwaWSMessage{
			Event:     "price",
			Symbol:    "XAU/USD",
			Price:     json.RawMessage(raw),
			Timestamp: time.Now().Unix(),
		}
		if err := scraper.handlePriceMessage(msg); err != nil {
			t.Fatalf("handlePriceMessage(%s): %v", raw, err)
		}

		if quote, queued := scraper.pendingQuotes["XAU/USD"]; queued {
			t.Errorf("price %s was queued for publishing as %v, want dropped", raw, quote.Price)
		}
	}
}

func TestHandlePriceMessageKeepsValidPrice(t *testing.T) {
	scraper := testCollectingScraper()

	msg := rwaWSMessage{
		Event:     "price",
		Symbol:    "XAU/USD",
		Price:     json.RawMessage(`4663.02946`),
		Timestamp: time.Now().Unix(),
	}
	if err := scraper.handlePriceMessage(msg); err != nil {
		t.Fatalf("handlePriceMessage: %v", err)
	}

	quote, queued := scraper.pendingQuotes["XAU/USD"]
	if !queued {
		t.Fatal("valid price was not queued for publishing")
	}
	if quote.Price != 4663.02946 {
		t.Errorf("queued price = %v, want 4663.02946", quote.Price)
	}
	if quote.Type != Commodities {
		t.Errorf("queued type = %v, want %v", quote.Type, Commodities)
	}
}

// The pending batch is a map keyed by symbol, so within one cooldown window a
// later tick replaces an earlier one. A zero must not be able to displace a
// price that was already collected.
func TestZeroPriceCannotOverwriteACollectedPrice(t *testing.T) {
	scraper := testCollectingScraper()

	good := rwaWSMessage{
		Event:     "price",
		Symbol:    "XAU/USD",
		Price:     json.RawMessage(`4663.02946`),
		Timestamp: time.Now().Unix(),
	}
	if err := scraper.handlePriceMessage(good); err != nil {
		t.Fatalf("handlePriceMessage(good): %v", err)
	}

	zero := good
	zero.Price = json.RawMessage(`0`)
	if err := scraper.handlePriceMessage(zero); err != nil {
		t.Fatalf("handlePriceMessage(zero): %v", err)
	}

	if got := scraper.pendingQuotes["XAU/USD"].Price; got != 4663.02946 {
		t.Errorf("pending price = %v, want the good price 4663.02946 to survive", got)
	}
}
