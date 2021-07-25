package strain

type Ints []int
type Lists [][]int
type Strings []string

func (list Ints) Keep(f func(int) bool) (out Ints) {
	for _, v := range list {
		if f(v) {
			out = append(out, v)
		}
	}

	return out
}

func (list Ints) Discard(f func(int) bool) (out Ints) {
	return list.Keep(func(n int) bool { return !f(n) })
}

func (list Lists) Keep(f func([]int) bool) (out Lists) {
	for _, v := range list {
		if f(v) {
			out = append(out, v)
		}
	}

	return out
}

func (list Strings) Keep(f func(string) bool) (out Strings) {
	for _, v := range list {
		if f(v) {
			out = append(out, v)
		}
	}

	return out
}
