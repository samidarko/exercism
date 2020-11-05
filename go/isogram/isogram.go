package isogram

import (
	"strings"
)

// IsIsogram takes a word and returns true if word is an isogram otherwise false
func IsIsogram(word string) bool {
	runeOccurred := map[rune]bool{}
	for _, r := range strings.ToLower(word) {
		if r == '-' || r == ' ' {
			continue
		}
		if runeOccurred[r] {
			return false
		}
		runeOccurred[r] = true
	}
	return true
}
