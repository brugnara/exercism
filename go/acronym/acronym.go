package acronym

import "strings"

// Abbreviate should abbreviate
func Abbreviate(s string) (ret string) {
	ar := strings.Split(strings.Replace(s, "-", " ", -1), " ")
	for _, word := range ar {
		chars := strings.Split(strings.Trim(word, " \n\t_-"), "")
		if len(chars) > 0 {
			ret += chars[0]
		}
	}
	ret = strings.ToUpper(ret)
	return
}
