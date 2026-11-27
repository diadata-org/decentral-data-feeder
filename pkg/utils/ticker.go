package utils

import (
	"math/rand/v2"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// POLL_JITTER_PERCENT_ENV sets how far each poll interval may stray from its nominal period,
	// as a percentage of that period. 0 disables jitter.
	POLL_JITTER_PERCENT_ENV     = "POLL_JITTER_PERCENT"
	defaultPollJitterPercent    = 10
	maxPollJitterPercentAllowed = 50
)

// JitteredTicker delivers ticks on C like time.Ticker, but each interval is the period plus a
// fresh uniform random offset in [-jitter, +jitter].
//
// Why: feeders that poll the same vendor every N seconds from one egress IP keep a fixed phase
// relative to each other. After a restart that phase is arbitrary, and if two of them land within
// a second of each other they stay there until the next restart. Against a vendor that refuses a
// second request per IP within about a second (Belo, 2026-11-27), the loser was refused on every
// poll for six hours. Re-drawing the offset on every tick means two pollers can still collide, but
// only on single ticks; they drift apart again instead of staying locked together.
//
// Like time.Ticker, a tick is dropped rather than queued when the reader has not taken the
// previous one, so a slow update cycle does not cause a burst of catch-up polls.
type JitteredTicker struct {
	C <-chan time.Time

	c      chan time.Time
	stop   chan struct{}
	once   sync.Once
	period time.Duration
	jitter time.Duration
}

// NewJitteredTicker returns a ticker whose intervals average period. jitter is clamped to
// [0, period/2] so an interval is never shorter than half the period.
func NewJitteredTicker(period, jitter time.Duration) *JitteredTicker {
	if period <= 0 {
		panic("utils.NewJitteredTicker: non-positive period")
	}
	if jitter < 0 {
		jitter = 0
	}
	if jitter > period/2 {
		jitter = period / 2
	}
	c := make(chan time.Time, 1)
	t := &JitteredTicker{C: c, c: c, stop: make(chan struct{}), period: period, jitter: jitter}
	go t.run()
	return t
}

// NewPollTicker returns a JitteredTicker for a vendor poll loop, with the jitter set by
// POLL_JITTER_PERCENT (default 10% of the period).
func NewPollTicker(period time.Duration) *JitteredTicker {
	return NewJitteredTicker(period, PollJitter(period))
}

// PollJitter is the jitter POLL_JITTER_PERCENT gives for period. An unparsable or out-of-range
// value logs an error and falls back to the default.
func PollJitter(period time.Duration) time.Duration {
	percent := defaultPollJitterPercent
	if raw := Getenv(POLL_JITTER_PERCENT_ENV, ""); raw != "" {
		p, err := strconv.Atoi(raw)
		if err != nil || p < 0 || p > maxPollJitterPercentAllowed {
			log.Errorf("%s=%q is not an integer in [0, %d]; using %d",
				POLL_JITTER_PERCENT_ENV, raw, maxPollJitterPercentAllowed, defaultPollJitterPercent)
		} else {
			percent = p
		}
	}
	return period * time.Duration(percent) / 100
}

// Stop ends the ticker. It does not close C, matching time.Ticker.
func (t *JitteredTicker) Stop() {
	t.once.Do(func() { close(t.stop) })
}

func (t *JitteredTicker) next() time.Duration {
	if t.jitter == 0 {
		return t.period
	}
	// Uniform in [-jitter, +jitter].
	offset := time.Duration(rand.Int64N(int64(2*t.jitter)+1)) - t.jitter
	return t.period + offset
}

func (t *JitteredTicker) run() {
	timer := time.NewTimer(t.next())
	defer timer.Stop()
	for {
		select {
		case <-t.stop:
			return
		case now := <-timer.C:
			select {
			case t.c <- now:
			default: // reader still busy with the previous tick: drop this one
			}
			timer.Reset(t.next())
		}
	}
}
