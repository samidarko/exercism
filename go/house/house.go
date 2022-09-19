package house

import (
	"fmt"
	"strings"
)

const totalVerses = 12

var subjects = []string{
	"the house that Jack built.",
	"the malt",
	"the rat",
	"the cat",
	"the dog",
	"the cow with the crumpled horn",
	"the maiden all forlorn",
	"the man all tattered and torn",
	"the priest all shaven and shorn",
	"the rooster that crowed in the morn",
	"the farmer sowing his corn",
	"the horse and the hound and the horn",
}

var verbs = []string{
	"is",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
	"",
}

func sentence(subjectIndex, verseIndex int) string {
	subject := subjects[subjectIndex]
	verb := verbs[subjectIndex]
	demonstrative := "that"
	if subjectIndex == verseIndex {
		// this is the first sentence
		demonstrative = "This"
		verb = verbs[0]
	}
	return fmt.Sprintf("%s %s %s", demonstrative, verb, subject)
}

func Verse(v int) string {
	var output strings.Builder

	if v == 1 {
		output.WriteString(sentence(0, 0))
		return output.String()
	}

	verseIndex := v - 1
	for subjectIndex := verseIndex; subjectIndex > 0; subjectIndex-- {
		output.WriteString(fmt.Sprintln(sentence(subjectIndex, verseIndex)))
	}

	output.WriteString(fmt.Sprintf("that lay in %s", subjects[0]))

	return output.String()
}

func Song() string {
	var output strings.Builder
	for v := 1; v < totalVerses; v++ {
		output.WriteString(fmt.Sprintf("%s\n\n", Verse(v)))
	}
	output.WriteString(Verse(totalVerses))
	return output.String()
}
