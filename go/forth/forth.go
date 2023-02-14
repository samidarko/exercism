package forth

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Stack []int

// Empty returns true if stack is empty
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// Push a new element to the stack
func (s *Stack) Push(element int) {
	*s = append(*s, element)
}

// Pop returns and remove the element on top of the stack
func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("empty Stack")
	}
	last := len(*s) - 1
	element := (*s)[last]
	*s = (*s)[:last]
	return element, nil
}

func (s *Stack) PopN(n int) ([]int, error) {
	values := make([]int, 0)
	for i := 0; i < n; i++ {
		value, err := s.Pop()
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

func Forth(input []string) ([]int, error) {
	stack := Stack{}
	words := map[string]string{}

	for _, instructions := range input {

		instructions = strings.ToLower(instructions)

		if word, value, ok := getWord(instructions); ok {

			if !isValidWord(word) {
				return nil, fmt.Errorf("invalid word %s", word)
			}

			for _, token := range strings.Split(value, " ") {
				if existingValue, found := words[token]; found {
					value = strings.ReplaceAll(value, token, existingValue)
				}
			}

			words[word] = value

			continue
		}

		instructions = checkInstructionsWords(instructions, words)

		for _, instruction := range strings.Split(instructions, " ") {
			switch instruction {
			case "+", "-", "*", "/":
				values, err := stack.PopN(2)
				if err != nil {
					return nil, err
				}
				b, a := values[0], values[1]
				result, err := operation(instruction, a, b)
				if err != nil {
					return nil, err
				}
				stack.Push(result)

			case "dup":
				a, err := stack.Pop()
				if err != nil {
					return nil, err
				}
				stack.Push(a)
				stack.Push(a)
			case "drop":
				_, err := stack.Pop()
				if err != nil {
					return nil, err
				}
			case "swap":
				values, err := stack.PopN(2)
				if err != nil {
					return nil, err
				}
				b, a := values[0], values[1]
				stack.Push(b)
				stack.Push(a)
			case "over":
				values, err := stack.PopN(2)
				if err != nil {
					return nil, err
				}
				b, a := values[0], values[1]
				stack.Push(a)
				stack.Push(b)
				stack.Push(a)

			default:
				value, err := strconv.Atoi(instruction)
				if err != nil {
					return nil, err
				}
				stack.Push(value)
			}

		}
	}

	return stack, nil
}

func operation(operator string, a, b int) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("integer divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator %s", operator)
	}
}

func getWord(instructions string) (string, string, bool) {
	if instructions[0] == ':' && instructions[len(instructions)-1] == ';' {
		instructions = instructions[2 : len(instructions)-2]
		parts := strings.SplitN(instructions, " ", 2)
		return parts[0], parts[1], true
	}
	return "", "", false
}

func checkInstructionsWords(instructions string, words map[string]string) string {
	for word, value := range words {
		instructions = strings.ReplaceAll(instructions, word, value)
	}
	return instructions
}

var isValidWord = regexp.MustCompile(`^[a-z-+/*]+$`).MatchString
