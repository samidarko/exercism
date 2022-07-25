package alphametics

import (
	"errors"
	"fmt"
	"strings"
)

func Solve(puzzle string) (map[string]int, error) {
	terms, sum, letters, valid := parse(puzzle)
	fmt.Println(terms)
	fmt.Println(sum)
	fmt.Println(letters)
	if !valid {
		return nil, errors.New("invalid puzzle")
	}
	panic("Please implement the Solve function")
}

func parse(puzzle string) ([]string, string, map[string]int, bool) {
	parts := strings.Split(puzzle, " == ")
	if len(parts) != 2 {
		return nil, "", nil, false
	}
	terms := strings.Split(parts[0], " + ")
	sum := parts[1]
	letters := map[string]int{}

	for _, word := range append(terms, sum) {
		for _, letter := range word {
			letters[string(letter)] = 0
		}
	}

	return strings.Split(parts[0], " + "), sum, letters, true
}
