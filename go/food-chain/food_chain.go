package foodchain

import (
	"fmt"
	"strings"
)

var animals = []string{
	"fly",
	"spider",
	"bird",
	"cat",
	"dog",
	"goat",
	"cow",
	"horse",
}

func getComment(animal string) string {
	switch animal {
	case "bird":
		return "How absurd to swallow a bird!"
	case "cat":
		return "Imagine that, to swallow a cat!"
	case "dog":
		return "What a hog, to swallow a dog!"
	case "goat":
		return "Just opened her throat and swallowed a goat!"
	case "cow":
		return "I don't know how she swallowed a cow!"
	default:
		return ""
	}
}

// Verse returns a verse
func Verse(v int) string {
	var output strings.Builder

	currentAnimal := animals[v-1]
	output.WriteString(fmt.Sprintf("I know an old lady who swallowed a %s.\n", currentAnimal))
	if currentAnimal == "horse" {
		output.WriteString(fmt.Sprint("She's dead, of course!"))
		return output.String()
	}
	if currentAnimal != "fly" {
		if comment := getComment(currentAnimal); comment != "" {
			output.WriteString(fmt.Sprintln(getComment(currentAnimal)))
		}
		if currentAnimal == "spider" {
			output.WriteString(fmt.Sprintf("It wriggled and jiggled and tickled inside her.\n"))
		}
		for i := v - 1; i >= 1; i-- {
			output.WriteString(fmt.Sprintf("She swallowed the %s to catch the %s", animals[i], animals[i-1]))
			if animals[i-1] == "spider" {
				output.WriteString(fmt.Sprint(" that wriggled and jiggled and tickled inside her"))
			}
			output.WriteString(fmt.Sprintln("."))
		}
	}
	output.WriteString("I don't know why she swallowed the fly. Perhaps she'll die.")
	return output.String()
}

// Verses returns verses within start and end inclusive
func Verses(start, end int) string {
	var output strings.Builder

	for s := start; s <= end-1; s++ {
		output.WriteString(fmt.Sprint(Verse(s), "\n\n"))
	}
	output.WriteString(Verse(end))

	return output.String()
}

// Song returns the song
func Song() string {
	return Verses(1, 8)
}
