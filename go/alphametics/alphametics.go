package alphametics

import (
	"fmt"
	"sort"
	"strings"
)

// Solve solves an alphametic puzzle by finding digit assignments that make the equation valid
// Time complexity: O(n!) where n is the number of unique letters
// Space complexity: O(n) for the recursion stack
func Solve(puzzle string) (map[string]int, error) {
	parts := strings.Split(puzzle, " == ")
	terms := strings.Split(parts[0], " + ")
	result := parts[1]

	// Collect unique letters and track first letters
	letters := make([]rune, 0)
	letterSet := make(map[rune]bool)
	firstLetters := make(map[rune]bool)

	for _, word := range append(terms, result) {
		firstLetters[rune(word[0])] = true
		for _, letter := range word {
			if !letterSet[letter] {
				letters = append(letters, letter)
				letterSet[letter] = true
			}
		}
	}

	// Sort letters for consistent ordering
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	// Generate permutations and check solutions
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	perms := permutations(digits, len(letters))

	for _, perm := range perms {
		// Check for leading zeros
		hasLeadingZero := false
		for letter := range firstLetters {
			if perm[getLetterIndex(letters, letter)] == 0 {
				hasLeadingZero = true
				break
			}
		}
		if hasLeadingZero {
			continue
		}

		// Create letter to digit mapping
		lettersMap := make(map[rune]int)
		for i, letter := range letters {
			lettersMap[letter] = perm[i]
		}

		// Check if the equation holds
		sum := 0
		for _, term := range terms {
			sum += wordVal(term, lettersMap)
		}

		if sum == wordVal(result, lettersMap) {
			resultMap := make(map[string]int)
			for letter, digit := range lettersMap {
				resultMap[string(letter)] = digit
			}
			return resultMap, nil
		}
	}

	return nil, fmt.Errorf("no solution")
}

// getLetterIndex returns the index of a letter in the sorted letters slice
func getLetterIndex(letters []rune, letter rune) int {
	for i, l := range letters {
		if l == letter {
			return i
		}
	}
	return -1
}

// wordVal calculates the numeric value of a word given the letter to digit mapping
func wordVal(word string, lettersMap map[rune]int) int {
	val := 0
	for _, ch := range word {
		val = val*10 + lettersMap[ch]
	}
	return val
}

// permutations generates all possible permutations of n digits from the given array
// Time complexity: O(n!) where n is the number of unique letters
// Space complexity: O(n) for the recursion stack
func permutations(arr []int, n int) [][]int {
	var result [][]int
	perm := make([]int, n)
	used := make([]bool, len(arr))

	var generate func(int)
	generate = func(pos int) {
		if pos == n {
			// Create a copy of the current permutation
			permCopy := make([]int, n)
			copy(permCopy, perm)
			result = append(result, permCopy)
			return
		}

		for i := 0; i < len(arr); i++ {
			if !used[i] {
				used[i] = true
				perm[pos] = arr[i]
				generate(pos + 1)
				used[i] = false
			}
		}
	}

	generate(0)
	return result
}
