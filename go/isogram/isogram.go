package isogram

import "strings"

// IsIsogram returns true if str is a isogram
func IsIsogram(str string) bool {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, "_", "")
	chars := strings.Split(str, "")
	seen := map[string]bool{}

	for _, char := range chars {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}
