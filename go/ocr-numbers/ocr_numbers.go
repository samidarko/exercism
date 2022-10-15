package ocr

import (
	"fmt"
	"strings"
)

// Recognize numbers on multiple lines
func Recognize(input string) (output []string) {
	lines := strings.Split(input, "\n")

	for i := 0; i < len(lines)-4; i += 4 {
		line := lines[i+1 : i+4]
		digits := ""
		for j := 0; j+3 <= len(line[0]); j += 3 {
			digit := fmt.Sprint(line[0][j:j+3], line[1][j:j+3], line[2][j:j+3])
			digits = fmt.Sprint(digits, recognizeDigit(digit))
		}
		output = append(output, digits)
	}

	return
}

func recognizeDigit(digit string) string {

	switch digit {
	case " _ | ||_|":
		return "0"
	case "     |  |":
		return "1"
	case " _  _||_ ":
		return "2"
	case " _  _| _|":
		return "3"
	case "   |_|  |":
		return "4"
	case " _ |_  _|":
		return "5"
	case " _ |_ |_|":
		return "6"
	case " _   |  |":
		return "7"
	case " _ |_||_|":
		return "8"
	case " _ |_| _|":
		return "9"
	default:
		return "?"
	}
}
