package strand

// "ACGTU" => [65 67 71 84 85]
var dnaRnaMapping = map[uint8]uint8{
	65: 85, // A => U
	67: 71, // C => G
	71: 67, // G => C
	84: 65, // T => A
}

// ToRNA Given a DNA strand, return its RNA complement (per RNA transcription)
func ToRNA(dna string) string {

	dnaBytes := []byte(dna)
	for i, _ := range dnaBytes {
		dnaBytes[i] = dnaRnaMapping[dnaBytes[i]]
	}
	return string(dnaBytes)
}
