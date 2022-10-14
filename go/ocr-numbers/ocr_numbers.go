package ocr

import (
	"fmt"
	"strings"
)

// Digit represents 3 x 3 'ascii' number
type Digit [3][3]rune

// Recognize numbers on multiple lines
func Recognize(input string) (output []string) {
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines)-4; i += 4 {
		line := lines[i+1 : i+4]
		digits := ""
		for j := 0; j+3 <= len(line[0]); j += 3 {
			digit := Digit{
				sliceToArray(line[0][j : j+3]),
				sliceToArray(line[1][j : j+3]),
				sliceToArray(line[2][j : j+3]),
			}
			digits = fmt.Sprint(digits, recognizeDigit(digit))
		}
		output = append(output, digits)
	}

	return
}

func sliceToArray(slice string) [3]rune {
	runes := []rune(slice)
	return [3]rune{runes[0], runes[1], runes[2]}
}

func recognizeDigit(digit Digit) string {

	switch {
	case digit[0][1] == ' ' && digit[2][0] == ' ' && digit[1][2] == '|':
		if digit[1][1] == '_' {
			return "4"
		}
		return "1"
	case digit[2][1] == ' ' && digit[2][0] == ' ' && digit[1][2] == '|':
		return "7"
	case digit[1][1] == ' ' && digit[0][1] == '_' && digit[2][0] == '|':
		return "0"
	case digit[1][0] == ' ' && digit[1][2] == '|':
		if digit[2][0] == '|' {
			return "2"
		}
		return "3"

	case digit[1][2] == ' ' && digit[1][0] == '|':
		if digit[2][0] == '|' {
			return "6"
		}
		return "5"

	case digit[0][1] == '_' && digit[1][1] == '_':
		if digit[2][0] == '|' {
			return "8"
		}
		return "9"

	default:
		return "?"
	}
}
