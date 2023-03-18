package scale

import (
	"fmt"
	"strings"
)

var flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}

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
func Scale(tonic, intervals string) []string {
	if len(intervals) == 0 {
		intervals = "mmmmmmmmmmm"
	}

	var scale []string

	switch tonic {
	case "C", "G", "D", "A", "E", "F#", "e", "b", "f#", "c#", "g#", "d#", "a":
		scale = sharps
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = flats
	default:
		panic(fmt.Sprintf("invalid tonic: %s", tonic))
	}

	index := tonicIndex(tonic, scale)
	notes := []string{scale[index]}

	for _, interval := range intervals {
		switch interval {
		case 'm':
			index += 1
		case 'M':
			index += 2
		case 'A':
			index += 3
		default:
			panic(fmt.Sprintf("invalid interval: %c", interval))
		}
		notes = append(notes, scale[index%len(scale)])
	}

	return notes
}
