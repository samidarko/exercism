package beer

import (
	"fmt"
	"strings"
)

// Song returns the song, all the verses from 99 to 0
func Song() string {
	song, _ := Verses(99, 0)
	return song
}

// Verses returns the verses between start and stop
func Verses(start, stop int) (string, error) {
	if start < stop {
		return "", fmt.Errorf("start less than stop")
	}
	var output strings.Builder
	for n := start; n >= stop; n-- {
		verse, err := Verse(n)
		if err != nil {
			return "", err
		}
		output.WriteString(fmt.Sprintln(verse))
	}
	return output.String(), nil
}

// Verse returns a verse
func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", fmt.Errorf("wrong verse")
	}
	if n == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}
	if n == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	}
	if n == 2 {
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	}
	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
}
