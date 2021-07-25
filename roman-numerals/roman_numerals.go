package romannumerals

import "fmt"

var translations = []struct {
	value int
	code  string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(n int) (o string, err error) {
	if 0 >= n || n > 3000 {
		return o, fmt.Errorf("invalid number: %d", n)
	}

	for _, t := range translations {
		for ; n-t.value >= 0; n -= t.value {
			o += t.code
		}
		if n == 0 {
			break
		}
	}

	return o, err
}
