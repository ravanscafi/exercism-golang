package phonenumber

import (
	"fmt"
	"regexp"
)

func Number(n string) (string, error) {
	nanp := regexp.MustCompile(`^\+?1`).ReplaceAllString(n, "")
	clean := regexp.MustCompile(`\D+`).ReplaceAllString(nanp, "")

	if !regexp.MustCompile(`^[2-9]\d{2}[2-9]\d{6}$`).MatchString(clean) {
		return "", fmt.Errorf("invalid number provided")
	}

	return clean, nil
}

func AreaCode(n string) (string, error) {
	c, err := Number(n)

	if err != nil {
		return "", err
	}

	return c[:3], nil
}

func Format(n string) (string, error) {
	c, err := Number(n)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%v) %v-%v", c[:3], c[3:6], c[6:]), nil
}
