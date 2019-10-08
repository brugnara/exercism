package raindrops

import "fmt"

// Convert will convert a number into rain string
func Convert(nr int) (ret string) {
	isDivisibleBy3 := nr%3 == 0
	isDivisibleBy5 := nr%5 == 0
	isDivisibleBy7 := nr%7 == 0

	if isDivisibleBy3 {
		ret += "Pling"
	}
	if isDivisibleBy5 {
		ret += "Plang"
	}
	if isDivisibleBy7 {
		ret += "Plong"
	}

	if !isDivisibleBy3 && !isDivisibleBy5 && !isDivisibleBy7 {
		ret = fmt.Sprintf("%d", nr)
	}

	return
}
