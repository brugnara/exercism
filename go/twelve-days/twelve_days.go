package twelve

import (
	"fmt"
	"strings"
)

// quick array to convert the *-th element to the corresponding string
var intToTh = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

// quick array with sentences
var sentences = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func initialSentence(line int) string {
	return fmt.Sprintf("On the %v day of Christmas my true love gave to me: ",
		intToTh[line-1])
}

// Verse returns the line-th sentence
func Verse(line int) (ret string) {
	var ar []string

	// sentences starts with the formatted line
	ret = initialSentence(line)

	// each line has the line-th sentence, plus all the previous joined together.
	for i := line; i > 1; i-- {
		ar = append(ar, sentences[i-1])
	}
	ret += strings.Join(ar, ", ")

	// if it is not the first line, append the "and"
	if len(ar) > 0 {
		ret += ", and "
	}

	// verse ends with a dot.
	ret += sentences[0] + "."

	return
}

// Song returns the whole song
func Song() (ret string) {
	len := len(sentences)
	ret = ""

	// iterate sentences in order to create the full Song
	for i := 0; i < len; i++ {
		ret += fmt.Sprintf("%v", Verse(i+1))

		// append \n only if it is not the last line
		if i != len-1 {
			ret += "\n"
		}
	}
	return
}
