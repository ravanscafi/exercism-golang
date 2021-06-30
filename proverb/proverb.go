// Package proverb generate proverbs.
package proverb

import "fmt"

// Proverb receive a list of rhyme and generate a proverb.
func Proverb(rhyme []string) []string {
	var strings []string
	max := len(rhyme)

	for i := range rhyme {
		if i == max-1 {
			strings = append(strings, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
			continue
		}
		strings = append(strings, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}

	return strings
}
