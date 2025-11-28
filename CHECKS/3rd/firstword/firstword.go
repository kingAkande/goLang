package main

import "fmt"

import "strings"

/**
Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').

A word is a sequence of characters delimited by spaces or by the start/end of the argument.

**/

func main() {

	fmt.Print(AKokoOro("hello there"))
	fmt.Print(AKokoOro(""))
	fmt.Print(AKokoOro("hello   .........  bye"))
}

// func  FirstWord  (arg string) string {

// 	word := strings.Fields(arg)

// 	res :="\n"

// 	if len(word)>0 {
// 	 return 	word[0] + res
// 	}

// 	return  res
// }

func AKokoOro(arg string) string {
	word := strings.Fields(arg)

	res := "\n"

	for len(word) > 0 {
		return word[0] + res
	}

	return res
}
