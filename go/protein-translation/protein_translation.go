package protein

import (
	"errors"
)

// ErrStop indicates a STOP codon was reached
var ErrStop = errors.New("ErrStop")

// ErrInvalidBase indicates an invalid codon
var ErrInvalidBase = errors.New("ErrInvalidBase")

// FromCodon takes a codon and returns a protein
func FromCodon(codon string) (string, error) {

	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA takes an rna a returns a list of proteins
func FromRNA(rna string) ([]string, error) {

	var proteins []string
	var processingError error

	for i := 0; i <= len(rna)-3; i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			break
		}
		if err == ErrInvalidBase {
			processingError = ErrInvalidBase
			break
		}
		proteins = append(proteins, protein)
	}
	return proteins, processingError
}
