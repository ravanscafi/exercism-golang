package raindrops

import "fmt"

func Convert(n int) string {
	r := ""

	if n % 3 == 0{
		r += "Pling"
	}

	if n % 5 == 0 {
		r += "Plang"
	}

	if n % 7 == 0 {
		r += "Plong"
	}

	if r == "" {
		return fmt.Sprint(n)
	}

	return r
}