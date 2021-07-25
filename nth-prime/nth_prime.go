package prime

import "math"

func Nth(n int) (int, bool) {
	i := 3

	if n == 1 {
		i++
	}

	for ; n > 1; i += 2 {
		if oddPrime(i) {
			n--
		}
	}

	return i - 2, n >= 1
}

func oddPrime(n int) bool {
	for i := 3; i <= int(math.Sqrt(float64(n))); i+= 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
