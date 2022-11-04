package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct Given a string of digits, calculate the largest product for a contiguous substring of digits of length n
func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span == 0 {
		return 1, nil
	}

	if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}

	if len(digits) < span {
		return -1, errors.New("span must be smaller than string length")
	}

	if isOnlyDigits(digits) == false {
		return -1, errors.New("digits input must only contain digits")
	}

	var largestSeriesProduct int64 = 0

	for i := 0; (i + span) <= len(digits); i++ {
		temp := seriesProduct(digits[i : i+span])
		if temp > largestSeriesProduct {
			largestSeriesProduct = temp
		}
	}

	return largestSeriesProduct, nil
}

func isOnlyDigits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) == false {
			return false
		}
	}
	return true
}

func seriesProduct(series string) (result int64) {
	result = int64(series[0] - '0')

	for _, r := range series[1:] {
		result = result * int64(r-'0')
	}
	return
}
