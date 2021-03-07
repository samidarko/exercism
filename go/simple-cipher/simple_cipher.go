package cipher

import (
	"log"
	"regexp"
	"strings"
)

type config struct {
	distances []int32
}

func sanitize(s string) string {
	reg, err := regexp.Compile("[^a-z]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(strings.ToLower(s), "")
}

func wrap(r rune) rune {
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

func (c config) Encode(s string) string {
	runes := []rune(sanitize(s))
	distanceIndex := 0
	distancesLength := len(c.distances)
	for i, r := range runes {
		r += c.distances[distanceIndex]
		distanceIndex++
		if distanceIndex == distancesLength {
			distanceIndex = 0
		}
		runes[i] = wrap(r)

	}
	return string(runes)
}

func (c config) Decode(s string) string {
	runes := []rune(s)
	distanceIndex := 0
	distancesLength := len(c.distances)
	for i, r := range runes {
		r -= c.distances[distanceIndex]
		distanceIndex++
		if distanceIndex == distancesLength {
			distanceIndex = 0
		}
		runes[i] = wrap(r)
	}
	return string(runes)
}

// NewCaesar returns a Caesar Cipher with a fixed shift distance of 3
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift returns a Shift Cipher with a flexible shift distance
func NewShift(shift int) Cipher {
	if shift == 0 || shift > 25 || shift < -25 {
		return nil
	}
	return config{distances: []int32{int32(shift)}}
}

// NewVigenere returns a Vigenère cipher
func NewVigenere(code string) Cipher {
	codeLength := len(code)
	if code == strings.Repeat("a", codeLength) || sanitize(code) != code {
		return nil
	}

	distances := make([]int32, codeLength)

	for i, r := range code {
		distances[i] = r - 'a'
	}

	return config{distances: distances}
}
