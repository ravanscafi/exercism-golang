package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("a and b must have the same length")
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
