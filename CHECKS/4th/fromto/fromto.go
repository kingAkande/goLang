package main

/*

Write a function that takes two integers and returns a string showing the range of numbers from the first to the second.

The numbers must be separated by a comma and a space.
If any of the arguments is bigger than 99 or less than 0, the function returns Invalid followed by a newline \n.
Prepend a 0 to any number that is less than 10.
Add a new line \n at the end of the string.

*/

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}


func FromTo(from int, to int) string {
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return  "invalid "
	}

	if from == to && from < 10{
		return  "0" + strconv.Itoa(from) + "\n"
	}

	return  ""
}