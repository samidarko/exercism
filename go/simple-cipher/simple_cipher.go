package cipher

import (
	"log"
	"regexp"
	"strings"
)

type Config struct {
	distance int32
}

func (c Config) Encode(s string) string {
	reg, err := regexp.Compile("[^a-z]+")
	if err != nil {
		log.Fatal(err)
	}
	runes := []rune(reg.ReplaceAllString(strings.ToLower(s), ""))
	for i, r := range runes {
		r += c.distance
		runes[i] = validateBoundaries(r)

	}
	return string(runes)
}

func (c Config) Decode(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		r -= c.distance
		runes[i] = validateBoundaries(r)
	}
	return string(runes)
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(shift int) Cipher {
	if shift == 0 || shift > 25 || shift < -25 {
		return nil
	}
	return Config{distance: int32(shift)}
}

//func NewVigenere(code string) Cipher {
//	return nil
//}

func validateBoundaries(r rune) rune {
	if r > 'z' {
		// if passed 'z' returns to 'a'
		r = 'a' + (r - ('z' + 1))
	}

	if r < 'a' {
		// if passed 'a' returns to 'z'
		r = 'z' - (('a' - 1) - r)
	}

	return r
}
