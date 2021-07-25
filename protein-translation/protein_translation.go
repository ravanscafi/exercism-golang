package protein

import (
	"errors"
)

var ErrStop = errors.New("stop codon")
var ErrInvalidBase = errors.New("invalid base")

var translations map[string]string = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(input string) (string, error) {
	value, ok := translations[input]

	if !ok {
		return "", ErrInvalidBase
	}

	if value == "STOP" {
		return "", ErrStop
	}

	return value, nil
}

func FromRNA(input string) ([]string, error) {
	var result []string

	for i := 0; i <= len(input)-3; i += 3 {
		value, err := FromCodon(input[i : i+3])

		if err == ErrStop {
			break
		}

		if err == ErrInvalidBase {
			return result, err
		}

		result = append(result, value)
	}

	return result, nil
}
