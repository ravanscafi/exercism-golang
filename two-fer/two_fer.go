// Package twofer has share functions.
package twofer

import "fmt"

// ShareWith shares something with given name or "you" otherwise.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
