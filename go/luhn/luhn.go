package luhn

import "strings"

func Valid(s string) bool {
	runes := []rune(strings.TrimSpace(s))
	length := len(runes)

	if length < 2 {
		return false
	}
	isSecondDigit := false
	digitsSum := 0

	var r rune

	for i := length - 1; i >= 0; i-- {
		r = runes[i]

		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			digit := int(r - '0')
			if isSecondDigit {
				doubledDigit := digit * 2
				if doubledDigit > 9 {
					digit = doubledDigit - 9
				} else {
					digit = doubledDigit
				}
				isSecondDigit = false
			} else {
				isSecondDigit = true
			}
			digitsSum += digit
		case ' ':
		default:
			return false

		}

	}
	return digitsSum%10 == 0
}
