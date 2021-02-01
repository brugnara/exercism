package scrabble

/*
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
*/
func runePoints(r rune) int {
	// a => A
	if r >= 97 {
		r -= 32
	}
	switch r {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	}
	return 0
}

// Score computes score for the given word
func Score(s string) (ret int) {
	for _, r := range s {
		ret += runePoints(r)
	}
	return
}
