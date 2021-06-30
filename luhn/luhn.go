package luhn

import (
	"strconv"
)

func Valid(c string) bool {
	r := 0
	j := 0

	for i := len(c) - 1; i >= 0; i-- {
		n := string(c[i])

		if n == " " {
			continue
		}

		d, _ := strconv.Atoi(n)

		if d == 0 && n != "0" {
			return false
		}

		if j%2 != 0 {
			d *= 2

			if d > 9 {
				d -= 9
			}
		}

		r += d
		j++
	}

	return (r > 0 || j > 1) && r%10 == 0
}
