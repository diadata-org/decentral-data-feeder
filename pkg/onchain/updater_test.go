package onchain

import (
	"math/big"
	"testing"
	"time"
)

func TestCollect(t *testing.T) {
	t0 := time.Date(2026, 9, 23, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name       string
		last       *publishedValue
		value      int64
		sourceTime time.Time
		want       bool
	}{
		{"first value with time", nil, 1, t0, true},
		{"first value without time", nil, 1, time.Time{}, true},
		{"value changed, newer time", &publishedValue{big.NewInt(1), t0}, 2, t0.Add(time.Second), true},
		{"value changed, same time", &publishedValue{big.NewInt(1), t0}, 2, t0, true},
		{"value changed, older time", &publishedValue{big.NewInt(1), t0}, 2, t0.Add(-time.Second), false},
		{"value unchanged, newer time", &publishedValue{big.NewInt(1), t0}, 1, t0.Add(time.Minute), false},
		{"no time, value changed", &publishedValue{big.NewInt(1), time.Time{}}, 2, time.Time{}, true},
		{"no time, value unchanged", &publishedValue{big.NewInt(1), time.Time{}}, 1, time.Time{}, false},
		{"no time, last had time, value changed", &publishedValue{big.NewInt(1), t0}, 2, time.Time{}, true},
		{"time, last had no time, value changed", &publishedValue{big.NewInt(1), time.Time{}}, 2, t0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys, values = nil, nil
			pending = make(map[string]publishedValue)
			lastPublished = make(map[string]publishedValue)
			if tt.last != nil {
				lastPublished["k"] = *tt.last
			}

			collect("k", big.NewInt(tt.value), tt.sourceTime)

			if got := len(keys) == 1; got != tt.want {
				t.Fatalf("collected = %v, want %v", got, tt.want)
			}
			if _, got := pending["k"]; got != tt.want {
				t.Fatalf("pending = %v, want %v", got, tt.want)
			}
			if tt.want && (values[0].Int64() != tt.value || !pending["k"].sourceTime.Equal(tt.sourceTime)) {
				t.Fatalf("got value %v time %v, want %v %v", values[0], pending["k"].sourceTime, tt.value, tt.sourceTime)
			}
		})
	}
}
