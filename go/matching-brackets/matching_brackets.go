package brackets

import (
	"fmt"
)

// Stack data structure
type Stack []interface{}

// Empty returns true if stack is empty
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// Push a new element to the stack
func (s *Stack) Push(element interface{}) {
	*s = append(*s, element)
}

// Pop returns and remove the element on top of the stack
func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return "", fmt.Errorf("empty Stack")
	}
	last := len(*s) - 1
	element := (*s)[last]
	*s = (*s)[:last]
	return element, nil
}

// Bracket returns true if all brackets are closed
func Bracket(input string) bool {
	stack := new(Stack)
	openingElement := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	for _, c := range input {
		switch c {
		case '[', '{', '(':
			stack.Push(c)
		case ']', '}', ')':
			last, err := stack.Pop()
			if err != nil || last.(rune) != openingElement[c] {
				return false
			}
		}

	}
	return stack.Empty()
}
