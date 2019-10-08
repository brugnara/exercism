package twofer

import "fmt"

// ShareWith will return a valid string
func ShareWith(name string) string {
	// if no name provided, sets name to "you"
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
