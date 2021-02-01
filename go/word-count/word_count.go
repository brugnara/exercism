package wordcount

import (
	"regexp"
	"strings"
)

// Frequency stores a map of seen words
type Frequency map[string]int

// WordCount returns a Frequency object with the counter for words
func WordCount(s string) Frequency {
	ret := Frequency{}
	for _, v := range strings.Fields(clean(s)) {
		// treats `'pippo'` as `pippo` but `can't` as `can't`
		ret[strings.Trim(v, "'")]++
	}
	return ret
}

// clean removes unnecessary words
func clean(s string) string {
	re := regexp.MustCompile(`[^a-z0-9'\s]`)
	return re.ReplaceAllString(strings.ToLower(s), " ")
}
