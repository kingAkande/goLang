package main

import "fmt"

/**

checknumber
Instructions
Write a function that takes a string as an argument and returns true if the string contains any number, otherwise return false.

*/

func main() {

	fmt.Println(CheckNumber("hello1"))
	fmt.Println(CheckNumber("hello"))

}

func CheckNumber(arg string) bool {

	if arg == "" {
		return false
	}

	for _, item := range arg {
		if item >= '0' && item <= '9' {
			return true
		}
	}
	return false
}
