package main

import "fmt"

/*
*

countalpha
Instructions
Write a function CountAlpha() that takes a string as an argument and returns the number of alphabetic characters in the string.
*/
func main() {

	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))

}

func CountAlpha(s string) int {
	count := 0
	for _, item := range s {
		if item >= 'a' && item <= 'z' || item >= 'A' && item <= 'Z' {
			count++
		}

	}
	return count
}
