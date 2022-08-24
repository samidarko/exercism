package wordsearch

import (
	"errors"
	"strings"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := map[string][2][2]int{}
	index := -1

	// horizontal search
	for _, word := range words {
		reversedWord := reverse(word)
		for i, row := range puzzle {
			// left to right
			index = strings.Index(row, word)
			if index > -1 {
				result[word] = [2][2]int{{index, i}, {index + len(word) - 1, i}}
			}
			// right to left
			index = strings.Index(row, reversedWord)
			if index > -1 {
				result[word] = [2][2]int{{index + len(word) - 1, i}, {index, i}}
			}
		}
	}

	// horizontal search
	for _, word := range words {
		reversedWord := reverse(word)
		for i, column := range getColumns(puzzle) {
			// top to bottom
			index = strings.Index(column, word)
			if index > -1 {
				result[word] = [2][2]int{{i, index}, {i, index + len(word) - 1}}
			}
			// bottom to top
			index = strings.Index(column, reversedWord)
			if index > -1 {
				result[word] = [2][2]int{{i, index + len(word) - 1}, {i, index}}
			}
		}
	}

	// diagonal search
	for _, word := range words {
		reversedWord := reverse(word)
		rowIndex := len(puzzle) - 1
		//colIndex := 0
		for _, diagonal := range getDiagonalsTopLeftBottomRight(puzzle) {
			// top left to bottom right
			index = strings.Index(diagonal, word)
			if index > -1 {
				result[word] = [2][2]int{{rowIndex, rowIndex + index}, {rowIndex + len(word) - 1, rowIndex + index + len(word) - 1}}
			}
			// bottom right to top left
			index = strings.Index(diagonal, reversedWord)
			if index > -1 {
				result[word] = [2][2]int{{index + len(word) - 1, rowIndex + index + len(word) - 1}, {index, rowIndex + index}}
			}
			if rowIndex > 0 {
				rowIndex--
			}
		}
	}

	if len(result) != len(words) {
		return result, errors.New("no result found")
	}
	return result, nil
}

func reverse(s string) string {
	output := make([]byte, len(s))
	for i, j := 0, len(s)-1; j >= 0; i, j = i+1, j-1 {
		output[i] = s[j]
	}
	return string(output)
}

func getColumns(puzzle []string) []string {
	columns := make([]string, len(puzzle[0]))
	for i := range columns {
		column := make([]uint8, len(puzzle))
		for j := range puzzle {
			column[j] = puzzle[j][i]
		}
		columns[i] = string(column)
	}
	return columns
}

func getDiagonalsTopLeftBottomRight(puzzle []string) []string {
	diagonals := make([]string, 0)

	for i, j := len(puzzle)-1, 0; j < len(puzzle); {
		diagonal := make([]uint8, 0)
		for x, y := i, j; x >= 0 && x < len(puzzle) && y >= 0 && y < len(puzzle[0]); x, y = x+1, y+1 {
			diagonal = append(diagonal, puzzle[x][y])
		}
		diagonals = append(diagonals, string(diagonal))
		if i > 0 {
			i--
		} else {
			j++
		}
	}

	return diagonals
}

//func getDiagonalsTopRightBottomLeft(puzzle []string) []string {
//	diagonals := make([]string, 0)
//
//	for i, j := len(puzzle)-1, 0; j < len(puzzle); {
//		diagonal := make([]uint8, 0)
//		for x, y := i, j; x >= 0 && x < len(puzzle) && y >= 0 && y < len(puzzle); x, y = x+1, y+1 {
//			diagonal = append(diagonal, puzzle[x][y])
//		}
//		diagonals = append(diagonals, string(diagonal))
//		if i > 0 {
//			i--
//		} else {
//			j++
//		}
//	}
//
//	return diagonals
//}
