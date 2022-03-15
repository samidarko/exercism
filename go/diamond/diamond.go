package diamond

import (
	"fmt"
	"strings"
)

// Gen diamond
func Gen(char byte) (string, error) {

	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("char out of range")
	}

	var output strings.Builder
	charPosition := char - byte('A') // nth char in alphabet (0 based)
	width := int(2*charPosition) + 1
	middle := int(charPosition)
	currentChar := byte('A')
	left, right := middle, middle

	for position := 0; position < width; position++ {
		row := []byte(strings.Repeat(" ", width))

		if position < middle {
			row[left], row[right] = currentChar, currentChar
			currentChar++
			left--
			right++
		} else {
			row[left], row[right] = currentChar, currentChar
			currentChar--
			left++
			right--
		}

		output.WriteString(fmt.Sprintln(string(row)))
	}

	return output.String(), nil
}
