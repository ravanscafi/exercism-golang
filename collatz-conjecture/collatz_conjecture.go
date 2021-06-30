package collatzconjecture

import "fmt"

func CollatzConjecture(number int) (int, error) {
	if number <= 0 {
		return -1, fmt.Errorf("number must be greater than 0")
	}

	return _run(number), nil
}

func _run(number int) int {
	if number == 1 {
		return 0
	}

	if number%2 == 0 {
		return _run(number/2) + 1
	}

	return _run(3*number+1) + 1
}
