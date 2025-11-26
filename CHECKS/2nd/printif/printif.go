package main

import "fmt"

/**
Write a function that takes a string as an argument and returns the letter G followed by a newline \n if the argument length is more or equal than 3,
otherwise returns Invalid Input followed by a newline \n.

If it's an empty string return G followed by a newline \n.

*/

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}

func PrintIf(str string) string {

	if len(str) >= 3 || str == "" {
		return "G\n"
	} else {
		return "Invalid Iput\n"
	}

}
