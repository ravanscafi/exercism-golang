// Package leap operates around leap years.
package leap

// IsLeapYear tells whether the given year is leap or not.
func IsLeapYear(y int) bool {
	return y % 4 == 0 && y % 100 != 0 || y % 400 == 0
}
