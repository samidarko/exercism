// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides Determine if a triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	switch {
	case !IsATriangle(a, b, c):
		k = NaT
	case a == b && b == c && c == a:
		k = Equ
	case a == b || b == c || c == a:
		k = Iso
	case a != b && b != c && c != a:
		k = Sca
	default:
		k = NaT

	}
	return k
}

func IsATriangle(a, b, c float64) bool {
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	return (a > 0 && b > 0 && c > 0) && (a+b >= c) && (b+c >= a) && (c+a >= b)
}
