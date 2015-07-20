package triangle

import "math"

// Kind is a kind of a triangle.
type Kind int

const (
	// NaT is not a triangle.
	NaT Kind = iota
	// Equ is a equilateral triangle.
	Equ
	// Iso is a isosceles triangle.
	Iso
	// Sca is a scalene triangle.
	Sca
)

// KindFromSides return a kind of a tringle from the lengths of 3 sides.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case isNeg(a, b, c) || isNaN(a, b, c) || isNotTriangle(a, b, c):
		return NaT
	case a == b && a == c:
		return Equ
	case a == b || b == c || c == a:
		return Iso
	default:
		return Sca
	}
}

func isNeg(a, b, c float64) bool {
	return a <= 0 || b <= 0 || c <= 0
}

func isNaN(a, b, c float64) bool {
	return math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)
}

func isNotTriangle(a, b, c float64) bool {
	return a+b <= c || a+c <= b || b+c <= a
}
