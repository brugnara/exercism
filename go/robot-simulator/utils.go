package robot

func minRU(a, b RU) RU {
	if a < b {
		return a
	}
	return b
}

func maxRU(a, b RU) RU {
	if a > b {
		return a
	}
	return b
}

func isValidCommand(cmd rune) bool {
	return cmd == 'A' || cmd == 'L' || cmd == 'R'
}
