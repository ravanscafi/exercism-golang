package reverse

func Reverse(s string) string {
	n := len(s)
	reversed := make([]rune, n)

	for _, r := range s {
		n--
	 	reversed[n] = r
	}

	return string(reversed[n:])
}
