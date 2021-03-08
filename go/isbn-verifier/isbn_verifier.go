package isbn

import (
	"strings"
	"unicode"
)

// IsValidISBN return true if isbn is valid
func IsValidISBN(isbn string) bool {

	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	factor := 10
	sum := 0

	for i, r := range isbn {
		if unicode.IsDigit(r) {
			sum += int(r-'0') * factor
			factor--
		} else if i == 9 && r == 'X' {
			sum += 10
		} else {
			return false
		}
	}

	return sum%11 == 0
}
