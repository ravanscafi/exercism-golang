package etl

import "strings"

func Transform(m map[int][]string) map[string]int {
	result := make(map[string]int)

	for value, letters := range m {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = value
		}
	}

	return result
}
