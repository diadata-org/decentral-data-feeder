package scraper

import (
	"encoding/json"
	"sort"
	"testing"
)

// Payload captured from production on 2026-08-19, when the XHKG/XFRA entitlement dropped
// and the partial failure sent the scraper into a reconnect loop.
const subscribeStatusWithFails = `{
	"event": "subscribe-status",
	"status": "warning",
	"success": [{"symbol": "AMZN", "exchange": "NASDAQ"}, {"symbol": "XAG/USD"}],
	"fails": [
		{"symbol": "1211", "message": "You are not authorized to access XHKG data."},
		{"symbol": "0700", "message": "You are not authorized to access XHKG data."},
		{"symbol": "HG1", "message": "You are not authorized to access XFRA data."}
	]
}`

func TestExtractFailedSymbols(t *testing.T) {
	var msg rwaWSMessage
	if err := json.Unmarshal([]byte(subscribeStatusWithFails), &msg); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	got := extractFailedSymbols(msg.Fails)
	sort.Strings(got)

	want := []string{"0700", "1211", "HG1"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestExtractFailedSymbolsEmpty(t *testing.T) {
	for name, raw := range map[string]string{
		"absent": `{"event":"subscribe-status","status":"ok"}`,
		"empty":  `{"event":"subscribe-status","status":"ok","fails":[]}`,
	} {
		var msg rwaWSMessage
		if err := json.Unmarshal([]byte(raw), &msg); err != nil {
			t.Fatalf("%s: unmarshal: %v", name, err)
		}
		if got := extractFailedSymbols(msg.Fails); len(got) != 0 {
			t.Errorf("%s: got %v, want none", name, got)
		}
	}
}

// A rejected symbol must be dropped from later subscribe calls, so the scraper stops
// re-requesting something the plan will never grant.
func TestFilterBlockedDropsRejectedSymbols(t *testing.T) {
	s := &RWAWSScraper{blockedSymbols: make(map[string]struct{})}

	all := []string{"1211", "0700", "AMZN", "XAG/USD", "HG1"}
	if got := s.filterBlocked(all); len(got) != len(all) {
		t.Fatalf("nothing blocked yet: got %v, want %v", got, all)
	}

	s.blockSymbols([]string{"1211", "0700", "HG1"})

	got := s.filterBlocked(all)
	want := []string{"AMZN", "XAG/USD"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}

	s.clearBlockedSymbols()
	if got := s.filterBlocked(all); len(got) != len(all) {
		t.Fatalf("after clear: got %v, want %v", got, all)
	}
}
