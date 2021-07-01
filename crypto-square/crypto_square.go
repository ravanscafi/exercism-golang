package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

func Encode(s string) string {
	s = clean(s)
	cols, rows := getRect(s)
	s = padStr(s, cols, rows)
	output := ""

	for c := 0; c < cols; {
		for r := 0; r < rows; r++ {
			output += string(s[c+r*cols])
		}

		if c++; c < cols {
			output += " "
		}
	}

	return output
}

func clean(s string) string {
	return regexp.MustCompile(`\W+`).ReplaceAllString(
		strings.ToLower(s),
		"",
	)
}

func getRect(s string) (int, int) {
	r := int(math.Round(math.Sqrt(float64(len(s)))))

	if r*r < len(s) {
		return r + 1, r
	}

	return r, r
}

func padStr(s string, cols, rows int) string {
	amount := cols * rows - len(s)

	for i := 0; i < amount; i++ {
		s += " "
	}

	return s
}
