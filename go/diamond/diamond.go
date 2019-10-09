package diamond

import (
	"errors"
	"fmt"
)

const lowerLimit = 65
const separator = " "

func getLine(nr, of int) (ret string) {
	// fmt.Printf("generating line #%d of %d\n", nr, of)
	lineLen := of*2 - 1
	spacesSides := of - nr
	spacesBetween := lineLen - spacesSides*2 - 2
	char := string(byte(nr + lowerLimit - 1))
	fmt.Printf("len: %d, sides: %d, between: %d\n",
		lineLen, spacesSides, spacesBetween)
	// adding element to the line
	// initial spaces
	for i := 0; i < spacesSides; i++ {
		ret += separator
	}
	// the char
	ret += char
	// spaces between, if any
	if spacesBetween > 0 {
		for i := 0; i < spacesBetween; i++ {
			ret += separator
		}
		// char, only if there are any spaces, otherwise this is the first or
		// the last line
		ret += char
	}
	// spaces after
	for i := 0; i < spacesSides; i++ {
		ret += separator
	}
	// and finally the \n
	ret += "\n"
	// fmt.Println(ret)
	return
}

// Gen generates the diamond starting with the given byte
func Gen(char byte) (ret string, err error) {
	if char < lowerLimit || char > 90 {
		return "", errors.New("invalid char error")
	}
	index := int(char - lowerLimit + 1)
	fmt.Printf("Got: %d = %v. Diamond will have h and l = %d\n",
		char, string(char), index)

	for i := 1; i <= index; i++ {
		ret += getLine(i, index)
	}
	// now doing in reverse for the bottom, just skip one line (index -1)
	for i := index - 1; i > 0; i-- {
		ret += getLine(i, index)
	}
	// printing the shape
	fmt.Println(ret)
	return
}
