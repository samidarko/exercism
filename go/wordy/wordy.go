package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	question = strings.ReplaceAll(question, "What is ", "")
	question = strings.ReplaceAll(question, " by", "")
	r := regexp.MustCompile("([+-]?\\d+|[a-zA-Z]+)")
	tokens := r.FindAllString(question, -1)

	var operation string
	var result int
	var err error

	if len(tokens) == 0 {
		return 0, false
	}

	result, err = strconv.Atoi(tokens[0])
	if err != nil {
		return 0, false
	}

	for i, element := range tokens[1:] {
		if i%2 == 0 {
			// should be an operation
			switch element {
			case "plus", "minus", "multiplied", "divided":
				operation = element
			default:
				return 0, false
			}
		} else {
			// should be a number
			number, err := strconv.Atoi(element)
			if err != nil {
				return 0, false
			}

			switch operation {
			case "multiplied":
				result *= number
			case "divided":
				result /= number
			case "minus":
				result -= number
			case "plus":
				result += number
			}
			operation = ""
		}
	}

	if operation != "" {
		return 0, false
	}
	return result, true
}
