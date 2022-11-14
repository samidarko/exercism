package scale

import (
	"fmt"
	"strings"
)

var flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

// parse is a sequence of intervals to absolute semitones above the tonic
func parse(s string) []int {
	if len(s) == 0 {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	}
	x := 0
	var notes []int
	for _, c := range s {
		notes = append(notes, x)
		x += map[rune]int{'m': 1, 'M': 2, 'A': 3}[c]
	}
	return notes
}

// tonicIndex the index of a tonic (as major or minor) in a list
func tonicIndex(tonic string, chromatic []string) int {
	for i := range chromatic {
		if strings.ToUpper(chromatic[i]) == strings.ToUpper(tonic) {
			return i
		}
	}
	panic(fmt.Sprintf("tonic %v not found", tonic))
}

// Scale returns the list of notes in scale given by a tonic and a set of intervals
func Scale(tonic, interval string) []string {
	notes := parse(interval)
	var out []string
	var chromatic []string
	switch tonic {
	case "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		chromatic = sharps
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		chromatic = flats
	default:
		chromatic = sharps
	}
	tonicI := tonicIndex(tonic, chromatic)
	for i := range notes {
		out = append(out, chromatic[(tonicI+notes[i])%12])
	}
	return out
}
