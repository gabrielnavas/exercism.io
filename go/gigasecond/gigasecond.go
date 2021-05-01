// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	seconds := time.Second * 1000000000
	duration := time.Duration(seconds)
	return t.Add(duration)
}
