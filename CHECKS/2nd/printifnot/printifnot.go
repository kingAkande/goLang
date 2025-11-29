package main

import "fmt"

/**
Write a function that takes a string and returns:

"G\n" if the string's length is less than 3 (including empty string).

"Invalid Input\n" otherwise.

*/

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}

func PrintIfNot(arg string) string {
	if len(arg) < 3 || arg == "" {
		return "G\n"
	} else {

		return "Invalid Input\n"
	}

}
