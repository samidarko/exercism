package atbash

import (
	"log"
	"regexp"
	"strings"
)

// Atbash substitution cipher
func Atbash(s string) string {
	var output strings.Builder
	alphabetMapping := map[rune]rune{}
	for r := 'a'; r <= 'z'; r++ {
		alphabetMapping[r] = 'z' - (r - 'a')
	}

	iteration := 0

	for _, r := range sanitize(s) {
		if iteration == 5 { // every 5 iteration insert a space
			output.WriteRune(' ')
			iteration = 0
		}
		if mapping, ok := alphabetMapping[r]; ok {
			output.WriteRune(mapping)
		} else {
			output.WriteRune(r)
		}
		iteration++
	}

	return output.String()
}

func sanitize(s string) string {
	reg, err := regexp.Compile("[^a-z1-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(strings.ToLower(s), "")
}
