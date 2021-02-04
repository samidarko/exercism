package cryptosquare

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"
)

func Encode(s string) string {
	runes := normalize(s)
	sLen := len(runes)
	sHalfLen := math.Sqrt(float64(sLen))

	r := int(sHalfLen)
	c := r + 1 // c >= r and c

	var output strings.Builder

	if float64(r) == sHalfLen {
		c = r
		r -= 1 // c >= r and c
	}

	fmt.Println("r", r, "c", c)

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			offset := i + j*c
			fmt.Println("offset", i, j, offset, sLen, offset%c)
			if offset < sLen {
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
	reg, err := regexp.Compile("[^a-z]+")
	if err != nil {
		log.Fatal(err)
	}
	return []rune(reg.ReplaceAllString(strings.ToLower(s), ""))
}
