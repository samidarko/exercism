// Package romannumerals provide functions to manipulate roman numerals
package romannumerals

import (
	"errors"
	"strings"
)

var arabicToRoman = map[int]string{
	1:    "I",
	2:    "II",
	3:    "III",
	4:    "IV",
	5:    "V",
	6:    "VI",
	7:    "VII",
	8:    "VIII",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

// ToRomanNumeral transform an arabic number to roman number
func ToRomanNumeral(arabic int) (string, error) {

	var roman string

	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("cannot convert 0, negative number or above 3000")
	}

	// remove the thousands
	arabic, roman = processMultipleOfTen(1000, arabic, roman)

	for _, value := range []int{900, 500, 400} {
		arabic, roman = processValue(value, arabic, roman)
	}

	// remove the hundreds
	arabic, roman = processMultipleOfTen(100, arabic, roman)

	for _, value := range []int{90, 50, 40} {
		arabic, roman = processValue(value, arabic, roman)
	}

	// remove the tens
	arabic, roman = processMultipleOfTen(10, arabic, roman)

	// at this stage arabic < 10
	roman += arabicToRoman[arabic]

	return roman, nil
}

func processMultipleOfTen(value, arabic int, roman string) (int, string) {

	numberOf := arabic / value

	return arabic - numberOf*value, roman + strings.Repeat(arabicToRoman[value], numberOf)
}

func processValue(value, arabic int, roman string) (int, string) {
	if arabic >= value {
		return arabic - value, roman + arabicToRoman[value]
	}
	return arabic, roman
}
