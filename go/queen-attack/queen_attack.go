package queenattack

import (
	"errors"
)

// CanQueenAttack return if two queens can attack each other.
// An error is returned if coords are invalid for a classic 8x8 chessboard
func CanQueenAttack(a, b string) (bool, error) {
	// basics checks
	if a == b || len(a) != 2 || len(b) != 2 {
		return false, errors.New("same position")
	}

	// Conversion to 0-7 values using int since we will need to use
	// a negative number, if any.
	x0 := int(a[0] - 97)
	y0 := int(a[1] - 49)
	x1 := int(b[0] - 97)
	y1 := int(b[1] - 49)

	// out of boundaries checks
	invalid0 := x0 < 0 || x0 > 7 || y0 < 0 || y0 > 7
	invalid1 := x1 < 0 || x1 > 7 || y1 < 0 || y1 > 7
	if invalid0 || invalid1 {
		return false, errors.New("invalid position")
	}
	// same row or same cell? easy peasy lemon squeezy!
	if x0 == x1 || y0 == y1 {
		return true, nil
	}

	// checking diagonals.
	return abs(x0-x1) == abs(y0-y1), nil
}

// abs for utility reasons.
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
