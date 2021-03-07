package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(s string, shift int) string {
	var output strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) {
			output.WriteRune(rotate(r, shift))
		} else {
			output.WriteRune(r)
		}
	}
	return output.String()
}

func rotate(r rune, shift int) rune {
	isUpper := unicode.IsUpper(r)
	r = unicode.ToLower(r)

	r += int32(shift)

	if r > 'z' {
		// if passed 'z' returns to 'a'
		r = 'a' + (r - ('z' + 1))
	}

	if isUpper {
		return unicode.ToUpper(r)
	}
	return r
}
