package alphametics

import (
	"errors"
	"fmt"
	"strings"
)

func pow(n, power int) int {
	if power == 0 {
		return 1
	}
	result := n
	for i := 2; i <= power; i++ {
		result *= n
	}
	return result
}

type Word struct {
	word    string
	letters map[string]int
}

func (w *Word) Value() int {
	result := 0
	position := len(w.letters)

	for _, l := range w.word {
		result += w.letters[string(l)] * pow(10, position)
		position--
	}

	return result
}

func Solve(puzzle string) (map[string]int, error) {
	terms, sum, letters, valid := parse(puzzle)
	if !valid {
		return nil, errors.New("invalid puzzle")
	}
	fmt.Println(terms)
	fmt.Println(sum)
	fmt.Println(letters)

	//data := map[string]int{"B": 9, "I": 1, "L": 0}

	// calculate all the letters

	return nil, nil
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
