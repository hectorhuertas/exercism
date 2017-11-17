// Package gigasecond implements a function to add a gigasecond to a date
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond adds a gigasecond.
func AddGigasecond(t time.Time) time.Time {
	giga := math.Pow10(9)
	gigasecond := time.Duration(giga) * time.Second
	return t.Add(gigasecond)
}
