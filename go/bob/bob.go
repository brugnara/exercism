// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob will reply to questions
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should reply accordingly to the given "remark"
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	// trimming the string
	remark = strings.Trim(remark, " \t\n\r")

	// bools for better reading
	hasLetters := strings.ContainsAny(remark, "qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM")
	isQuestion := strings.HasSuffix(remark, "?")
	isShout := hasLetters && (strings.ToUpper(remark) == remark)

	// empty sentence
	if remark == "" {
		return "Fine. Be that way!"
	}

	// self documented
	if isQuestion {
		if isShout {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if hasLetters && isShout {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
