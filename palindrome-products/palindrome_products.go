package palindrome

import (
	"fmt"
	"math"
	"strconv"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	var n int
	pmin.Product = math.MaxInt32
	pmax.Product = math.MinInt32

	if fmin > fmax {
		err = fmt.Errorf("fmin > fmax: %d %d", fmin, fmax)
		return
	}

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			n = i * j

			if !Palindrom(n) {
				continue
			}

			if n > pmax.Product {
				pmax = Product{n, [][2]int{{i, j}}}
			} else if n == pmax.Product {
				pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
			}

			if n < pmin.Product {
				pmin = Product{n, [][2]int{{i, j}}}
			} else if n == pmin.Product {
				pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
			}
		}
	}

	if pmax.Product == math.MinInt32 {
		err = fmt.Errorf("no palindromes: %d %d", fmin, fmax)
	}

	return
}

func Palindrom(n int) bool {
	s := strconv.Itoa(int(math.Abs(float64(n))))
	l := len(s)

	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}

	return true
}
