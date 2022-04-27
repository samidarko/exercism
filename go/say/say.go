package say

import (
	"fmt"
	"strings"
)

// Say spell out a number n in English
func Say(n int64) (string, bool) {
	if n < 0 || n >= 1000000000000 {
		return "", false
	}

	output := make([]string, 0)
	numbers := []int64{1000000000, 1000000, 1000, 100}

	for _, i := range numbers {
		count := n / i
		if count > 0 {
			say, _ := Say(count)
			output = append(output, fmt.Sprintf("%s %s", say, numberName[i]))
			n -= count * i
			if n == 0 {
				n = -1
				break
			}
		}
	}

	if n > 20 {
		tens := n / 10
		firstDigit, secondDigit := tens*10, n-tens*10
		output = append(output, fmt.Sprintf("%s-%s", numberName[firstDigit], numberName[secondDigit]))
		n = -1
	}
	if n >= 0 {
		output = append(output, numberName[n])
	}
	return strings.Join(output, " "), true
}

var numberName = map[int64]string{
	0:          "zero",
	1:          "one",
	2:          "two",
	3:          "three",
	4:          "four",
	5:          "five",
	6:          "six",
	7:          "seven",
	8:          "eight",
	9:          "nine",
	10:         "ten",
	11:         "eleven",
	12:         "twelve",
	13:         "thirteen",
	14:         "fourteen",
	15:         "fifteen",
	16:         "sixteen",
	17:         "seventeen",
	18:         "eighteen",
	19:         "nineteen",
	20:         "twenty",
	30:         "thirty",
	40:         "forty",
	50:         "fifty",
	60:         "sixty",
	70:         "seventy",
	80:         "eighty",
	90:         "ninety",
	100:        "hundred",
	1000:       "thousand",
	1000000:    "million",
	1000000000: "billion",
}
