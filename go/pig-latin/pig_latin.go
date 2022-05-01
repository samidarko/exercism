package piglatin

import "strings"

func isVowel(r rune) bool {
	vowels := "aeiouy"
	for _, v := range vowels {
		if v == r {
			return true
		}
	}
	return false
}

const separator = " "

// Sentence returns a sentence in Pig Latin
func Sentence(sentence string) string {
	words := make([]string, 0)

	for _, word := range strings.Split(sentence, separator) {
		words = append(words, Word(word))
	}

	return strings.Join(words, separator)
}

// Word returns a word in Pig Latin
func Word(sentence string) string {
	var output strings.Builder

	switch {
	case len(sentence) == 2 && sentence[1] == 'y':
		output.WriteRune('y')
		output.WriteRune(rune(sentence[0]))
	case (isVowel(rune(sentence[0])) && sentence[:2] != "ye") || sentence[:2] == "xr":
		output.WriteString(sentence)
	case sentence[1:3] == "qu":
		output.WriteString(sentence[3:])
		output.WriteRune(rune(sentence[0]))
		output.WriteString("qu")
	case sentence[:3] == "thr" || sentence[:3] == "sch":
		output.WriteString(sentence[3:])
		output.WriteString(sentence[:3])
	case sentence[:2] == "ch" || sentence[:2] == "qu" || sentence[:2] == "th" || sentence[:2] == "rh":
		output.WriteString(sentence[2:])
		output.WriteString(sentence[:2])
	default:
		output.WriteString(sentence[1:])
		output.WriteRune(rune(sentence[0]))
	}
	output.WriteString("ay")
	return output.String()
}
