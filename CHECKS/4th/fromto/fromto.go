package main

import (
	"fmt"
)

/*
Write a function that takes two integers and returns a string showing the range of numbers from the first to the second.

The numbers must be separated by a comma and a space.
If any of the arguments is bigger than 99 or less than 0, the function returns Invalid followed by a newline \n.
Prepend a 0 to any number that is less than 10.
Add a new line \n at the end of the string.


**/

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}


func FromTo(from int, to int) string {
	// Invalid input check
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	// If both numbers are the same
	if from == to {
		return fmt.Sprintf("%02d\n", from)
	}

	result := ""

	// Count up (from < to)
	if from < to {
		for i := from; i <= to; i++ {
			result += fmt.Sprintf("%02d", i)
			if i != to {
				result += ", "
			}
		}
	}

	// Count down (from > to)
	if from > to {
		for i := from; i >= to; i-- {
			result += fmt.Sprintf("%02d", i)
			if i != to {
				result += ", "
			}
		}
	}

	return result + "\n"
}


