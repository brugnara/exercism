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
