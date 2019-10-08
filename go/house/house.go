package house

import (
	"fmt"
	"strings"
)

var verses = []string{
	"that lay in|the house that Jack built",
	"that ate|the malt",
	"that killed|the rat",
	"that worried|the cat",
	"that tossed|the dog",
	"that milked|the cow with the crumpled horn",
	"that kissed|the maiden all forlorn",
	"that married|the man all tattered and torn",
	"that woke|the priest all shaven and shorn",
	"that kept|the rooster that crowed in the morn",
	"that belonged to|the farmer sowing his corn",
	"|the horse and the hound and the horn",
}

func capitalize(line string) string {
	tmp := strings.Split(line, "")
	tmp[0] = strings.ToUpper(tmp[0])
	return strings.Join(tmp, "")
}

// Verse will return the  partial song starting from given line
func Verse(line int) (verse string) {
	// tmp := strings.Split(verses[line-1], "|")

	// first element always starts with "this is "
	// verse = capitalize("this is " + tmp[1])

	isFirst := true

	for i := line; i > 0; i-- {
		tmp := strings.Split(verses[i-1], "|")
		prefix := "this is"

		if !isFirst {
			prefix = tmp[0]
		}
		isFirst = false

		verse += fmt.Sprintf("%v %v", prefix, tmp[1])

		if i > 1 {
			verse += "\n"
		}
	}

	return capitalize(verse) + "."
}

// Song will return the whole song
func Song() (ret string) {
	ret = ""
	len := len(verses)

	// iterating verses
	for i := 0; i < len; i++ {
		ret += Verse(i + 1)

		// append separator only if not last item
		if i != len-1 {
			ret += "\n\n"
		}
	}
	return
}
