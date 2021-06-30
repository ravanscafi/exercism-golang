package diffsquares

func SquareOfSum(n int) int {
	r := 0

	for i := 1; i <= n; i++ {
		r += i
	}

	return r * r
}

func SumOfSquares(n int) int {
	r := 0

	for i := 1; i <= n; i++ {
		r += i*i
	}

	return r
}

func Difference(n int) int {
	a := 0
	b := 0

	for i := 1; i <= n; i++ {
		a += i
		b += i*i
	}

	return a*a - b
}
