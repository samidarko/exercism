package cryptosquare

import (
	"log"
	"math"
	"regexp"
	"strings"
)

func Encode(s string) string {
	runes := normalize(s)
	sqrt := math.Sqrt(float64(len(runes)))
	r := int(math.Round(sqrt))
	c := r

	if float64(r) < sqrt {
		c = r + 1
	}

	var output strings.Builder

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			offset := i + j*c
			if offset < len(runes) {
				output.WriteRune(runes[offset])
			} else {
				output.WriteRune(' ')
			}
		}
		if i < c-1 {
			output.WriteRune(' ')
		}
	}

	return output.String()
}

func normalize(s string) []rune {
	reg, err := regexp.Compile("[^a-z1-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return []rune(reg.ReplaceAllString(strings.ToLower(s), ""))
}
