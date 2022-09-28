package luhn

import (
	"strings"
	"unicode"
)

// Valid given a number determine whether or not it is valid per the Luhn formula
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")

	if len(s) < 2 { // input smaller than 2 cannot be valid
		return false
	}

	isSecondDigit := len(s)%2 == 0 // determines if first character is second digit or not
	digitsSum := 0

	for _, r := range s {

		if !unicode.IsDigit(r) { // non digit characters makes input invalid
			return false
		}

		digit := int(r - '0') // gives real digit value not rune value
		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		isSecondDigit = !isSecondDigit
		digitsSum += digit
	}
	return digitsSum%10 == 0
}
