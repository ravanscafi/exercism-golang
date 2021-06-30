// Package triangle handle triangle operations.
package triangle

import (
	"math"
)

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides determines which kind of triangle from its sides.
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 || math.IsNaN(a+b+c) || math.IsInf(a*b*c,0) || a+b < c || a+c < b || b+c < a   {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
