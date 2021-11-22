package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	question = strings.ReplaceAll(question, "What is ", "")
	question = strings.ReplaceAll(question, "multiplied by", "multipliedby")
	question = strings.ReplaceAll(question, "divided by", "dividedby")
	r := regexp.MustCompile("([+-]?\\d+|[a-zA-Z]+)")
	tokens := r.FindAllString(question, -1)

	var operation string
	var result int
	var err error

	if len(tokens) >= 1 {
		result, err = strconv.Atoi(tokens[0])
		if err != nil {
			return 0, false
		}
	} else {
		return 0, false
	}

	for i, element := range tokens[1:] {
		if i%2 == 0 {
			// should be an operation
			switch element {
			case "plus", "minus", "multipliedby", "dividedby":
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

			result = calculate(result, number, operation)
			operation = ""
		}
	}

	if operation != "" {
		return 0, false
	}
	return result, true
}

func calculate(a, b int, operation string) int {
	switch operation {
	case "multipliedby":
		return a * b
	case "dividedby":
		return a / b
	case "minus":
		return a - b
	default:
		// plus
		return a + b
	}
}
