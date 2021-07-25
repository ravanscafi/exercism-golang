package anagram

import "strings"

func Detect(s string, list []string) (out []string) {
	s = strings.ToLower(s)

	for _, c := range list {
		if IsAnagram(s, c) {
			out = append(out, c)
		}
	}

	return out
}

func IsAnagram(s, c string) bool  {
	c = strings.ToLower(c)

	if s == c || len(c) != len(s) {
		return false
	}

	for _, r := range c {
		i := strings.IndexRune(s, r)
		if i == -1 {
			return false
		}
		s = s[:i] + s[i+1:]
	}

	return true
}