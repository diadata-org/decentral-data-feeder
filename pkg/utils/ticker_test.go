package utils

import (
	"testing"
	"time"
)

func TestJitteredTickerIntervalsStayInBounds(t *testing.T) {
	period, jitter := 40*time.Millisecond, 10*time.Millisecond
	tk := NewJitteredTicker(period, jitter)
	defer tk.Stop()

	for i := 0; i < 200; i++ {
		d := tk.next()
		if d < period-jitter || d > period+jitter {
			t.Fatalf("interval %v outside [%v, %v]", d, period-jitter, period+jitter)
		}
	}
}

func TestJitteredTickerIntervalsVary(t *testing.T) {
	tk := NewJitteredTicker(60*time.Second, 6*time.Second)
	defer tk.Stop()

	seen := map[time.Duration]bool{}
	for i := 0; i < 50; i++ {
		seen[tk.next()] = true
	}
	if len(seen) < 10 {
		t.Fatalf("only %d distinct intervals in 50 draws; jitter is not being applied", len(seen))
	}
}

func TestJitteredTickerZeroJitterIsFixed(t *testing.T) {
	tk := NewJitteredTicker(50*time.Millisecond, 0)
	defer tk.Stop()
	for i := 0; i < 20; i++ {
		if d := tk.next(); d != 50*time.Millisecond {
			t.Fatalf("zero jitter gave %v", d)
		}
	}
}

func TestJitteredTickerClampsJitterToHalfPeriod(t *testing.T) {
	tk := NewJitteredTicker(100*time.Millisecond, time.Second)
	defer tk.Stop()
	if tk.jitter != 50*time.Millisecond {
		t.Fatalf("jitter %v, want clamped to 50ms", tk.jitter)
	}
}

func TestJitteredTickerTicksAndStops(t *testing.T) {
	tk := NewJitteredTicker(20*time.Millisecond, 5*time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case <-tk.C:
		case <-time.After(time.Second):
			t.Fatalf("tick %d never arrived", i)
		}
	}
	tk.Stop()
	tk.Stop() // idempotent
	// Drain a tick that may already be buffered, then expect silence.
	select {
	case <-tk.C:
	default:
	}
	select {
	case <-tk.C:
		t.Fatal("tick after Stop")
	case <-time.After(100 * time.Millisecond):
	}
}

func TestJitteredTickerDropsTicksForSlowReader(t *testing.T) {
	tk := NewJitteredTicker(20*time.Millisecond, 0)
	defer tk.Stop()
	time.Sleep(150 * time.Millisecond) // ~7 periods with nobody reading
	<-tk.C
	// Checked immediately: a queued backlog would be sitting here; the next real tick is ~20ms off.
	select {
	case <-tk.C:
		t.Fatal("ticks were queued; a slow reader would get a burst of catch-up polls")
	default:
	}
}

func TestPollJitter(t *testing.T) {
	cases := []struct {
		env  string
		want time.Duration
	}{
		{"", 6 * time.Second},    // default 10%
		{"0", 0},                 // disabled
		{"25", 15 * time.Second}, // explicit
		{"abc", 6 * time.Second}, // unparsable -> default
		{"80", 6 * time.Second},  // above the cap -> default
		{"-5", 6 * time.Second},  // negative -> default
	}
	for _, c := range cases {
		t.Setenv(POLL_JITTER_PERCENT_ENV, c.env)
		if got := PollJitter(60 * time.Second); got != c.want {
			t.Errorf("POLL_JITTER_PERCENT=%q: got %v, want %v", c.env, got, c.want)
		}
	}
}
