// Package etl extract, transform and load
package etl

import "strings"

// Transform transform a map of score => letters into letter => score
func Transform(m map[int][]string) map[string]int {
	result := map[string]int{}

	for score, letters := range m {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}

	return result
}
