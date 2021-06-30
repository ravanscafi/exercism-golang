package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

func (list IntList) Length() int {
	length := 0

	for range list {
		length++
	}

	return length
}

func (list IntList) Reverse() IntList {
	l := list.Length()
	result := make(IntList, l)

	for i, v := range list {
		result[l-1-i] = v
	}

	return result
}

func (list IntList) Append(other IntList) IntList {
	result := make(IntList, list.Length()+other.Length())

	copy(result, list)
	copy(result[list.Length():], other)

	return result
}

func (list IntList) Concat(args []IntList) IntList {
	var result = list

	for _, v := range args {
		result = result.Append(v)
	}

	return result
}

func (list IntList) Map(f unaryFunc) IntList {
	result := make(IntList, list.Length())

	for i, v := range list {
		result[i] = f(v)
	}

	return result
}

func (list IntList) Filter(f predFunc) IntList {
	intermediate := make(IntList, list.Length())
	i := 0

	for _, v := range list {
		if f(v) {
			intermediate[i] = v
			i++
		}
	}

	result := make(IntList, i)
	copy(result[0:i], intermediate)

	return result
}

func (list IntList) Foldl(f binFunc, acc int) int {
	for _, v := range list {
		acc = f(acc, v)
	}

	return acc
}

func (list IntList) Foldr(f binFunc, acc int) int {
	for _, v := range list.Reverse() {
		acc = f(v, acc)
	}

	return acc
}
