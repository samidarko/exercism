package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	song := make([]string, 0)
	for i := 0; i < takeDown; i++ {
		bottleCount := startBottles - i
		firstSentence := fmt.Sprintf("%s green bottle%s hanging on the wall,", digitToString(bottleCount, true), plural(bottleCount))
		verse := []string{
			firstSentence,
			firstSentence,
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green bottle%s hanging on the wall.", digitToString(bottleCount-1, false), plural(bottleCount-1)),
			"",
		}
		song = append(song, verse...)
	}

	return song[:len(song)-1]
}

func plural(i int) string {
	if i == 1 {
		return ""
	}
	return "s"
}

func digitToString(i int, capitalize bool) string {
	digit := ""
	switch i {
	case 10:
		digit = "ten"
	case 9:
		digit = "nine"
	case 8:
		digit = "eight"
	case 7:
		digit = "seven"
	case 6:
		digit = "six"
	case 5:
		digit = "five"
	case 4:
		digit = "four"
	case 3:
		digit = "three"
	case 2:
		digit = "two"
	case 1:
		digit = "one"
	case 0:
		digit = "no"
	default:
		return ""
	}

	if capitalize {
		return strings.ToUpper(digit[0:1]) + digit[1:]
	}

	return digit
}
