// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, " ", "  ")
	s = strings.ReplaceAll(s, "_", "")
	chrs := ""
	for _, word := range strings.Split(s, " ") {
		if len(word) > 0 {
			word = strings.ToUpper(word)
			chrs += string(word[0])
		}
	}
	return chrs
}
