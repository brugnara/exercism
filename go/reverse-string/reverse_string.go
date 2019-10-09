package reverse

import "strings"

// Reverse reverse an input string
func Reverse(str string) (ret string) {
	// transforms a string into an array of chars
	tmp := strings.Split(str, "")
	len := len(tmp)

	// reading the array from the end and adding to ret each char.
	for i := len; i > 0; i-- {
		ret += tmp[i-1]
	}

	return
}
