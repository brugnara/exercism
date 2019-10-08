package triangle

import "math"

// Kind is the type we returns in KindFromSides
type Kind int

// Const
const (
	NaT = 1 // not a triangle
	Equ = 2 // equilateral
	Iso = 3 // isosceles
	Sca = 4 // scalene
)

func isValid(a float64) bool {
	if a <= 0 {
		return false
	}
	if math.IsNaN(a) {
		return false
	}
	if math.IsInf(a, 1) {
		return false
	}
	if math.IsInf(a, -1) {
		return false
	}
	return true
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) (k Kind) {
	// invalid side

	// invalid side len
	if (a+b < c) || (a+c < b) || (b+c < a) {
		return NaT
	}

	if !isValid(a) || !isValid(b) || !isValid(c) {
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
