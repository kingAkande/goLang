package main

import "strings"

/**
Write a function LastWord that takes a string and returns its last word followed by a \n.

A word is a section of string delimited by spaces or by the start/end of the string.


*/

func LastWord(arg string) string {
	word := strings.Fields(arg)

	if len(word) > 0 {
		return word[len(word) - 1] + "\n"
	}

	return "\n"
}
