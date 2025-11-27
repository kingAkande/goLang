package main

import "fmt"

import "strings"

/**
Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').

A word is a sequence of characters delimited by spaces or by the start/end of the argument.

**/

func main() {

	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {

}
