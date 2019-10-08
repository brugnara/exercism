package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

// Valid checks is str is valid
func Valid(str string) bool {
	fmt.Printf("Checking string: %s\n", str)
	str = strings.Replace(str, " ", "", -1)
	chars := strings.Split(str, "")
	l := len(chars)

	if l < 2 {
		return false
	}

	// no extra chars allowed
	if strings.ContainsAny(str,
		"qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM") {
		return false
	}

	sum := 0
	// reading the chars array starting from the end. We must check that i is even
	// or odd since
	for i := 0; i < l; i++ {
		index := l - i - 1
		value := chars[index]
		nr, _ := strconv.Atoi(value)
		fmt.Println(nr)
		fmt.Printf("Index: %v\n", index)
		if i%2 == 1 {
			fmt.Printf("Doubling: %v at index #%v\n", nr, i)
			nr *= 2
			if nr > 9 {
				fmt.Println("More than 9, subtracting 9")
				nr -= 9
			}
		}
		fmt.Printf("Adding: %v\n", nr)
		sum += nr
	}

	valid := sum%10 == 0
	fmt.Printf("Sum: %d. Valid? %v\n", sum, valid)

	return valid
}
