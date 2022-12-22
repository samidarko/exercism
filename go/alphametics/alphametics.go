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

func Combinations(L []int, r int) [][]int {
	if r == 1 {
		//Convert every item in L to List and
		//Append it to List of List
		temp := make([][]int, 0)
		for _, rr := range L {
			t := make([]int, 0)
			t = append(t, rr)
			temp = append(temp, [][]int{t}...)
		}
		return temp
	} else {
		res := make([][]int, 0)
		for i := 0; i < len(L); i++ {
			//Take only elements till i
			// remember we do not care about position
			perms := make([]int, 0)
			perms = append(perms, L[:i]...)
			for _, x := range Combinations(perms, r-1) {
				t := append(x, L[i])
				res = append(res, [][]int{t}...)
			}
		}
		return res
	}
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

		for _, term := range terms {
			word, err := NewWord(term, data)
			if err != nil {
				continue
			}
			total += word.Value()
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
