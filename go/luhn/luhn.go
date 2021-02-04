package luhn

import (
	"strings"
	"unicode"
)

// Valid given a number determine whether or not it is valid per the Luhn formula
func Valid(s string) bool {
	runes := []rune(strings.ReplaceAll(s, " ", ""))
	length := len(runes)

	if length < 2 {
		return false
	}
	isSecondDigit := false
	digitsSum := 0

	var r rune

	for i := length - 1; i >= 0; i-- {
		r = runes[i]

		if !unicode.IsDigit(r) {
			return false
		}

		digit := int(r - '0')
		if isSecondDigit {
			doubledDigit := digit * 2
			if doubledDigit > 9 {
				digit = doubledDigit - 9
			} else {
				digit = doubledDigit
			}
		}

		isSecondDigit = !isSecondDigit
		digitsSum += digit
	}
	return digitsSum%10 == 0
}
