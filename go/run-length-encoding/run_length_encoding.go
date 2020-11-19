package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode run length encode
func RunLengthEncode(s string) (encoding string) {

	var character rune
	var repetition int

	for _, c := range s {
		if c != character {
			encoding += Encode(character, repetition)
			character, repetition = c, 1
		} else {
			repetition++
		}
	}

	encoding += Encode(character, repetition)

	return

}

// Encode the character
func Encode(character rune, repetition int) string {

	switch repetition {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%c", character)
	default:
		// more than 1 repetition
		return fmt.Sprintf("%d%c", repetition, character)
	}

}

// RunLengthDecode run length decode
func RunLengthDecode(s string) (decoding string) {

	var number []rune
	for _, c := range s {
		if unicode.IsDigit(c) {
			number = append(number, c)
		} else {
			decoding += Decode(c, number)
			number = []rune{}
		}
	}

	return
}

// Decode the code
func Decode(character rune, number []rune) string {
	if len(number) == 0 {
		return fmt.Sprintf("%c", character)
	}
	if repetition, err := strconv.Atoi(string(number)); err == nil {
		return strings.Repeat(string(character), repetition)
	}
	return ""
}
