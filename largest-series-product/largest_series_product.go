package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int, error) {
	result := 0

	if 0 > span || span > len(digits) {
		return 0, errors.New("wrong span")
	}

	for i := 0; i <= len(digits) - span; i++ {
		partial := 1

		for j := 0; j < span; j++ {
			val, err := strconv.Atoi(string(digits[i+j]))
			if err != nil {
				return 0, err
			}
			partial *= val
			if partial == 0 {
				i += j
				break
			}
		}

		if partial > result {
			result = partial
		}
	}

	return result, nil
}
