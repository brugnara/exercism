package ocr

import (
	"fmt"
	"log"
	"strings"
)

const (
	zero = `
 _
| |
|_|
   `
	one = `

  |
  |
   `
	two = `
 _
 _|
|_
   `
	three = `
 _
 _|
 _|
   `
	four = `

|_|
  |
   `
	five = `
 _
|_
 _|
   `
	six = `
 _
|_
|_|
   `
	seven = `
 _
  |
  |
   `
	eight = `
 _
|_|
|_|
   `
	nine = `
 _
|_|
 _|
   `
	charLen    = 3
	lineHeight = 5
)

var mapper = map[string]string{
	cleanTrailingSpaces(zero):  "0",
	cleanTrailingSpaces(one):   "1",
	cleanTrailingSpaces(two):   "2",
	cleanTrailingSpaces(three): "3",
	cleanTrailingSpaces(four):  "4",
	cleanTrailingSpaces(five):  "5",
	cleanTrailingSpaces(six):   "6",
	cleanTrailingSpaces(seven): "7",
	cleanTrailingSpaces(eight): "8",
	cleanTrailingSpaces(nine):  "9",
}

// recognizeDigit not used.
func recognizeDigit(str []string) int {
	fmt.Println(str)
	return 1
}

// a line is 4 lines (4x \n, the 4th is empty)
func extractCharsFromSingleLine(str string) (digits []string) {
	// split the line into matrix of 3x4. Each matrix will contains the numbers
	lines := strings.Split(str, "\n")
	l := len(lines)
	// preventing errors
	if l > lineHeight {
		log.Fatalf("Maximum 4 lines allowed here, got: %v. Please split if multiple lines", l)
	}
	lineLen := 0
	// find longest lineLen
	for _, line := range lines {
		l := len(line)
		if l > lineLen {
			lineLen = l
		}
	}

	fmt.Printf("LineLen: %d\n", lineLen)

	// seeking cursor every charLen in order to read an OCR char per time
	for cursor := 0; cursor < lineLen; cursor += charLen {
		fmt.Printf("Cursor: %d\n", cursor)
		digit := "\n"
		// a char is splitted into lineHeight lines, so we are joining them
		// TODO: this produces a very verbose output:
		for y := 0; y < lineHeight; y++ {
			fmt.Printf("y: %d\n", y)
			fmt.Printf("Working on line: %v\n", lines[y])
			chars := strings.Split(lines[y], "")
			charsInCurrentLine := len(chars)
			fmt.Printf("line has %d chars\n", charsInCurrentLine)
			if charsInCurrentLine == 0 {
				fmt.Println("empty line")
				continue
			}
			lineSplit := ""
			// adding a char per time to the lineSplit var
			for x := 0; x < charLen; x++ {
				index := x + cursor
				fmt.Printf("x: %d. Reading index: %d\n", x, index)
				lineSplit += chars[index]
				fmt.Printf("Line is now: '%s'\n", lineSplit)
			}
			if y < lineHeight-1 {
				lineSplit += "\n"
			}
			fmt.Printf("Linesplit: '%s'\n", lineSplit)
			// joining by \n the lineSplit will produce a complete OCR number.
			digit += cleanTrailingSpaces(lineSplit)
		}
		// This allow to clearly see the char read
		fmt.Print("Digit >>>")
		fmt.Print(printHelper(digit))
		fmt.Print("<<<\n")

		// producing an array of digits
		digits = append(digits, digit)
	}

	return
}

// printHelper will replace \s with a dot. This simple method will allow to
// easily debug the output.
func printHelper(str string) string {
	return fmt.Sprint(strings.ReplaceAll(str, " ", "."))
}

func cleanTrailingSpaces(str string) string {
	tmp := strings.Split(str, "\n")
	for i, line := range tmp {
		tmp[i] = strings.TrimRight(line, " ")
	}
	return strings.Join(tmp, "\n")
}

func convert(str string) string {
	fmt.Printf("Wanting to convert string: '%s'\n", printHelper(str))
	str = cleanTrailingSpaces(str)
	fmt.Printf("Trailed string: '%s'\n", printHelper(str))
	ret := mapper[str]
	if ret == "" {
		return "?"
	}
	return ret
}

// Recognize a string
func Recognize(str string) (ret []string) {
	fmt.Println(str)
	// - iterate through each "line", composed by 4 lines each (\n as separator)
	lines := strings.Split(str, "\n")
	totalLen := len(lines)
	fmt.Printf("Total len is: '%d'\n", totalLen)

	// we are reading event the line "before". This will cause to read a whole
	// empty line on next iterations. A bigLine is composed by lineHeight
	// elements but we are going to consider the line separating each bigLines.
	// In order to do so, we are incrementing bigLine by lineHeight -1 and then
	// skip the first line, if bigLine > 0. This is not a monkey-patch, it is the
	// way I found to skip the empty line used as separator. The same line that
	// is considered empty for the bottom of a number and the same considered
	// as the initial \n char.
	for bigLine := 0; bigLine < totalLen; bigLine += lineHeight - 1 {
		fmt.Printf("current bigLine# is: '%d'\n", bigLine)
		if lineHeight+bigLine > totalLen {
			fmt.Println("out of index, safe to ignore and skip this")
			continue
		}
		currentLines := []string{}
		startingIndex := 0

		// if not on the first "bigLine", the "next" first line will be an empty
		// one. Since we don't need that, we are skipping that line. Since we need
		// 5 lines in order to have a "complete" bigLine, we are adding a "" element
		// that will generate a "\n" only when Joined creating "currentLine".
		// The "currentLine" will be then processed by "extractCharsFromSingleLine"
		// with no effort.
		if bigLine > 0 {
			startingIndex = 1
			currentLines = append(currentLines, "")
		}
		for i := startingIndex; i < lineHeight; i++ {
			currentLines = append(currentLines, lines[i+bigLine])
		}
		currentLine := strings.Join(currentLines, "\n")
		// DONT!
		// currentLine = cleanTrailingSpaces(currentLine)
		// working on currentLine
		fmt.Printf("Current line is: '%s'\n", printHelper(currentLine))

		// - strip each line in array of numbers
		currentLineNumber := ""
		numbers := extractCharsFromSingleLine(currentLine)
		fmt.Printf("We have %d numbers to convert here:\n%v\n",
			len(numbers), strings.Join(numbers, ", "))
		for _, nr := range numbers {
			number := convert(nr)
			fmt.Printf("Converted number is %s\n", number)
			currentLineNumber += fmt.Sprintf("%v", number)
		}
		fmt.Printf("Current line number is: '%v'\n", currentLineNumber)
		//
		ret = append(ret, currentLineNumber)
	}

	return
}
