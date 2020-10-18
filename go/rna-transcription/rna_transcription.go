package strand

var dnaRnaMapping = map[rune]rune{
	'A': 'U',
	'C': 'G',
	'G': 'C',
	'T': 'A',
}

// ToRNA Given a DNA strand, return its RNA complement (per RNA transcription)
func ToRNA(dna string) string {

	dnaBytes := make([]rune, len(dna))
	for i, r := range dna {
		dnaBytes[i] = dnaRnaMapping[r]
	}
	return string(dnaBytes)
}
