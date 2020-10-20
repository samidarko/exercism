package isogram

import (
	"strings"
)

// IsIsogram takes a word and returns true if word is an isogram otherwise false
func IsIsogram(word string) bool {
	runesOccurences := map[rune]int{}
	for _, r := range strings.ToLower(word) {
		if r == '-' || r == ' ' {
			continue
		}
		runesOccurrence := runesOccurences[r]
		if runesOccurrence == 1 {
			return false
		}
		runesOccurences[r] = runesOccurrence + 1
	}
	return true
}
