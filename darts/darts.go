package darts

func Score(x, y float64) int {
	dist := x*x + y*y

	if dist <= 1 {
		return 10
	}

	if dist <= 25 {
		return 5
	}

	if dist <= 100 {
		return 1
	}

	return 0
}
