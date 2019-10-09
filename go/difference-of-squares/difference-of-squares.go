package diffsquares

import "math"

// SquareOfSum junk
func SquareOfSum(nr int) int {
	value := 0
	for i := 0; i < nr; i++ {
		value += i + 1
	}
	return int(math.Pow(float64(value), 2.0))
}

// SumOfSquares junk
func SumOfSquares(nr int) int {
	value := 0.0
	for i := 0; i < nr; i++ {
		value += math.Pow(float64(i+1), 2.0)
	}
	return int(value)
}

// Difference junk
func Difference(nr int) int {
	return SquareOfSum(nr) - SumOfSquares(nr)
}
