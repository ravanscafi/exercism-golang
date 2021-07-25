package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) (result []string) {
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, UnsafeFirst(n, s[i:]))
	}

	return result
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

// First returns the first substring of s with length n, if n is not bigger than s length
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}

	return UnsafeFirst(n, s), true
}

