package utils

import "time"

func CheckFreshness(timestamp time.Time, freshness time.Duration) (bool, time.Duration) {
	age := time.Since(timestamp)
	over := age - freshness
	if over <= 0 {
		return true, 0
	}
	return false, over
}
