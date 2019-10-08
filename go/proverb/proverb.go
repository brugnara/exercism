package proverb

import "fmt"

func basePhrase(w1, w2 string) string {
	return fmt.Sprintf("For want of a %v the %v was lost.", w1, w2)
}

func lastPhrase(w string) string {
	return fmt.Sprintf("And all for the want of a %v.", w)
}

// Proverb should provide a proverb
func Proverb(rhyme []string) (ret []string) {
	len := len(rhyme)
	if len < 1 {
		return
	}
	for i := 1; i < len; i++ {
		ret = append(ret, basePhrase(rhyme[i-1], rhyme[i]))
	}
	ret = append(ret, lastPhrase(rhyme[0]))
	return
}
