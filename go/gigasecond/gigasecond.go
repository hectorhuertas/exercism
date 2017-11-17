// Package gigasecond implements a function to add a gigasecond to a date
package gigasecond

import "time"

// AddGigasecond adds a gigasecond.
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(1e9) * time.Second
	return t.Add(gigasecond)
}
