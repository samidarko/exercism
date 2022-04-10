package railfence

import (
	"strings"
)

func Encode(message string, rails int) string {
	message = strings.ReplaceAll(message, " ", "")
	var output strings.Builder
	grid := getGrid(message, rails)

	for i := 0; i < rails; i++ {
		for j := range message {
			r := grid[i][j]
			if r > 0 {
				output.WriteRune(r)
			}
		}
	}

	return output.String()
}

func Decode(message string, rails int) string {
	grid := getGrid(message, rails)

	// update grid with message (currently grid is created with encrypted message)
	for rail, charIndex := 0, 0; rail < rails; rail++ {
		for i := range message {
			if grid[rail][i] > 0 {
				grid[rail][i] = rune(message[charIndex])
				charIndex++
			}
		}
	}

	// read the grid
	var output strings.Builder
	rail, inc := 0, 1

	for i := range message {
		r := grid[rail][i]
		if r > 0 {
			output.WriteRune(r)
		}
		nextRail := rail + inc
		if nextRail < 0 || nextRail >= rails {
			inc *= -1
		}
		rail += inc
	}

	return output.String()
}

func getGrid(message string, rails int) [][]rune {
	grid := make([][]rune, rails)
	for i := range grid {
		grid[i] = make([]rune, len(message))
	}

	rail, inc := 0, 1
	for i, r := range message {
		grid[rail][i] = r
		nextRail := rail + inc
		if nextRail < 0 || nextRail >= rails {
			inc *= -1
		}
		rail += inc
	}

	return grid
}

//func displayGrid(grid [][]rune) string {
//	var output strings.Builder
//
//	for _, runes := range grid {
//		for _, r := range runes {
//
//			if r > 0 {
//				output.WriteRune(r)
//			} else {
//				output.WriteRune('.')
//			}
//
//		}
//		output.WriteRune('\n')
//	}
//
//	return output.String()
//}
