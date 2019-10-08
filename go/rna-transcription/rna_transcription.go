package strand

import "strings"

var mapper = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA converter
func ToRNA(dna string) (ret string) {
	tmp := strings.Split(dna, "")

	for _, item := range tmp {
		ret += mapper[item]
	}

	return
}
