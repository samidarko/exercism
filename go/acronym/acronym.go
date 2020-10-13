// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

var r = regexp.MustCompile(`[\s-]`)

// Abbreviate Convert a phrase to its acronym.
func Abbreviate(s string) string {

	words := r.Split(s, -1)

	var acronym string

	for _, word := range words {

		word = strings.Trim(word, "_")
		if word != "" {
			acronym += strings.ToUpper(word[0:1])
		}

	}

	return acronym
}
