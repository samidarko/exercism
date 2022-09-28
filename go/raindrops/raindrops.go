// Package raindrops contains functions to convert raindrop sounds
package raindrops

import "strconv"

// Convert a number into a string that contains raindrop sounds
func Convert(i int) string {
	result := ""

	if i%3 == 0 {
		result += "Pling"
	}

	if i%5 == 0 {
		result += "Plang"
	}

	if i%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(i)
	}
	return result
}
