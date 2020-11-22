package wordcount

import (
	"regexp"
	"strings"
)

// Frequency type
type Frequency = map[string]int

// WordCount count words for a given sentence and return a Frequency
func WordCount(sentence string) Frequency {

	var frequency = make(Frequency)
	r := regexp.MustCompile(`[\n:!@$%^&:,.]`)

	splitSentence := strings.Split(r.ReplaceAllString(strings.ToLower(sentence), " "), " ")

	for _, word := range splitSentence {
		word = strings.Trim(word, "'")
		if word != "" {
			frequency[word] = frequency[word] + 1
		}
	}

	return frequency
}
