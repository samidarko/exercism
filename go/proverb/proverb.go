// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

// Proverb For want of a horseshoe nail, a kingdom was lost, or so the saying goes.
func Proverb(rhyme []string) []string {
	lenRhyme := len(rhyme)
	if lenRhyme == 0 {
		return []string{}
	}

	text := make([]string, lenRhyme)
	for i := 0; i < lenRhyme-1; i++ {
		text[i] = ForWantOfAThingTheOtherThingWasLost(rhyme[i], rhyme[i+1])
	}

	text[lenRhyme-1] = AndAllForTheWantOfAThing(rhyme[0])

	return text

}

func ForWantOfAThingTheOtherThingWasLost(thing, otherThing string) string {
	return "For want of a " + thing + " the " + otherThing + " was lost."
}

func AndAllForTheWantOfAThing(thing string) string {
	return "And all for the want of a " + thing + "."
}
