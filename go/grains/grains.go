package grains

import "errors"

// Square get the value of a square
func Square(nr int) (ret uint64, err error) {

	if nr < 1 || nr > 64 {
		return 0, errors.New("invalid input")
	}

	ret = 1

	for i := 0; i < nr-1; i++ {
		ret *= 2
	}

	return
}

// Total gives the total
func Total() (total uint64) {
	// adding the amount on each Square in order to get the grand total
	for i := 1; i < 65; i++ {
		value, _ := Square(i)
		total += value
	}
	return
}
