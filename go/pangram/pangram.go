package pangram

import (
	"strings"
)

// IsPangram determine if a sentence is a pangram
func IsPangram(s string) bool {

	occurrenceMap := getOccurrenceMap()

	for _, c := range strings.ToLower(s) {
		if _, ok := occurrenceMap[c]; ok {
			occurrenceMap[c] = true
		}
	}

	return allOccurred(occurrenceMap)
}

// returns a map where each key is a letter from 'a' to 'z' and the value a boolean set as false
func getOccurrenceMap() map[rune]bool {
	occurrenceMap := map[rune]bool{}
	var i int32
	const c = 'a'
	for i < 26 {
		occurrenceMap[c+i] = false
		i++
	}
	return occurrenceMap
}

// test all occurrences occurred at least once
func allOccurred(occurrenceMap map[rune]bool) bool {
	for _, occurred := range occurrenceMap {
		if occurred == false {
			return false
		}
	}
	return true
}
