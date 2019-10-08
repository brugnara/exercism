package hamming

import (
	"errors"
	"strings"
)

// Distance will return the distance between DNA
func Distance(a, b string) (distance int, err error) {
	lenA := len(a)
	lenB := len(b)

	// we are not able to calculate if len are different
	if lenA != lenB {
		err = errors.New("a")
		return
	}

	// splitting strings into arrays of "chars"
	ar1 := strings.Split(a, "")
	ar2 := strings.Split(b, "")

	// let's iterate the first array and check with the second one
	for i, item1 := range ar1 {
		if item1 != ar2[i] {
			distance++
		}
	}

	return
}
