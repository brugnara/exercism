package accumulate

// Accumulate performs a func over the given slice
// and returns a new slice with the edits
func Accumulate(sx []string, fn func(string) string) []string {
	ret := make([]string, len(sx))

	for i, s := range sx {
		ret[i] = fn(s)
	}

	return ret
}
