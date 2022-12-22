package alphametics

import (
	"errors"
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
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

func NewWord(word string, letters map[string]int) (*Word, error) {
	if letters[string(word[0])] == 0 {
		return nil, fmt.Errorf("cannot have 0 for first letter")
	}
	return &Word{
		word:    word,
		letters: letters,
	}, nil
}

func (w *Word) Value() int {
	value := 0
	position := len(w.word) - 1

	for _, l := range w.word {
		value += w.letters[string(l)] * pow(10, position)
		position--
	}

	return value
}

func Solve(puzzle string) (map[string]int, error) {
	terms, sum, letters, valid := parse(puzzle)
	if !valid {
		return nil, errors.New("invalid puzzle")
	}

	for _, permutation := range combin.Permutations(10, len(letters)) {
		data := map[string]int{}

		for i, l := range letters {
			data[string(l)] = permutation[i]
		}

		total := 0

		words := map[string]int{}

		for _, term := range terms {

			value, ok := words[term]

			//if ok {
			//	fmt.Println("found")
			//}

			if !ok {
				word, err := NewWord(term, data)
				if err != nil {
					continue
				}
				value = word.Value()
				words[term] = value
			}

			total += value
		}

		result, err := NewWord(sum, data)
		if err != nil {
			continue
		}

		resultValue := result.Value()

		if total == resultValue {
			return data, nil
		}

	}

	return nil, fmt.Errorf("no solution")
}

func parse(puzzle string) ([]string, string, string, bool) {
	parts := strings.Split(puzzle, " == ")
	if len(parts) != 2 {
		return nil, "", "", false
	}
	terms := strings.Split(parts[0], " + ")
	sum := parts[1]
	letters := make([]rune, 0)
	letterSet := map[rune]bool{}

	for _, word := range append(terms, sum) {
		for _, letter := range word {
			if !letterSet[letter] {
				letters = append(letters, letter)
			}
			letterSet[letter] = true
		}
	}

	return strings.Split(parts[0], " + "), sum, string(letters), true
}
