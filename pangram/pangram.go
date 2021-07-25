package pangram

import "strings"

func IsPangram(s string) bool {
	s = strings.ToLower(s)

	for i := 'a'; i <= 'z'; i++ {
		if !strings.Contains(s, string(i)) {
			return false
		}
	}

	return true
}
