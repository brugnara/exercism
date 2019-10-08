package beer

import (
	"errors"
	"fmt"
)

// Verses returns beer verses
func Verses(from, to int) (ret string, err error) {
	if from > 99 || to < 0 || from <= to {
		return "", errors.New("ko")
	}
	for i := from; i >= to; i-- {
		tmp, _ := Verse(i)
		ret += tmp + "\n"
	}
	return
}

// Verse returns beer verses
func Verse(nr int) (string, error) {
	if nr < 0 || nr > 99 {
		return "", errors.New("s")
	}
	if nr == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	}
	if nr == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}
	word := "bottles"
	if nr == 2 {
		word = "bottle"
	}
	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d %v of beer on the wall.\n",
		nr, nr, nr-1, word), nil
}

// Song a
func Song() string {
	tmp, _ := Verses(99, 0)
	return tmp
}
