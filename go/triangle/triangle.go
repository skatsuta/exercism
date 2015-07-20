package triangle

import (
	"math"
	"sort"
)

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
	// sort 3 sides
	s := []float64{a, b, c}
	sort.Float64Slice(s).Sort()

	// check to contain NaN or infinity
	for _, e := range s {
		if math.IsInf(e, 0) || math.IsNaN(e) {
			return NaT
		}
	}

	// check to satisfy the triangle inequality
	if s[0]+s[1] <= s[2] {
		return NaT
	}

	if s[0] != s[1] && s[1] != s[2] {
		return Sca
	}

	if s[0] == s[2] {
		return Equ
	}

	return Iso
}
