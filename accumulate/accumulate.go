package accumulate

func Accumulate(l []string, f func(string) string) (a []string) {
	a = make([]string, len(l))

	for i, item := range l {
		a[i] = f(item)
	}

	return
}
