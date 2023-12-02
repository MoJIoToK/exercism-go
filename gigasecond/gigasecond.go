// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

// import path for the time package from the standard library
import "time"

const GIGASECUND = 1000000000

// AddGigasecond return time after 1 gigasecond from the entry time.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * GIGASECUND)
	return t
}
