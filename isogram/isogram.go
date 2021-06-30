package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(w string) bool {
	w = strings.ToUpper(w)

	for i, r := range w {
		if unicode.IsLetter(r) && strings.ContainsRune(w[i+1:], r) {
			return false
		}
	}

	return true
}
