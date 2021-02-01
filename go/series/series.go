package series

// All computes all
func All(n int, s string) []string {
	ret := []string{}
	for i := 0; i < len(s)-n+1; i++ {
		ret = append(ret, s[i:i+n])
	}
	return ret
}

// UnsafeFirst is a trap
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return s, false
	}
	return UnsafeFirst(n, s), true
}
