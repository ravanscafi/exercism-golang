// Package gigasecond handle gigasecond operations.
package gigasecond

import "time"

// AddGigasecond add a gigasecond to given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1_000_000_000 * time.Second)
}
