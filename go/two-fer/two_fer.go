// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import (
	"strings"
)

func ShareWith(name string) string {
	var phrase string = "One for X, one for me."
	if len(name) == 0 {
		return strings.ReplaceAll(phrase, "X", "you")
	}
	return strings.ReplaceAll(phrase, "X", name)
}
