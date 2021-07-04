package grains

import "fmt"

const MaxSq = 64

func Square(n int) (uint64, error) {
	if 1 > n || n > MaxSq {
		return 0, fmt.Errorf("square number must be between 1 and %d, %d given", MaxSq, n)
	}

	return 1 << (n - 1), nil
}

func Total() (t uint64) {
	for i := 1; i <= MaxSq; i++ {
		v, _ := Square(i)
		t += v
	}

	return
}
