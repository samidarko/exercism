// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

var isAlpha = regexp.MustCompile(`[A-Za-z]+`).MatchString

// Hey should reply as Bob would
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case isAlpha(remark) && remark == strings.ToUpper(remark) && strings.Contains(remark, "?"):
		return "Calm down, I know what I'm doing!"
	case isAlpha(remark) && remark == strings.ToUpper(remark):
		return "Whoa, chill out!"
	case remark != "" && remark[len(remark)-1] == '?':
		return "Sure."
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}
