package main

import "strconv"

/*
Write a function that takes two integers and returns a string showing the range of numbers from the first to the second.

The numbers must be separated by a comma and a space.
If any of the arguments is bigger than 99 or less than 0, the function returns Invalid followed by a newline \n.
Prepend a 0 to any number that is less than 10.
Add a new line \n at the end of the string.


**/

func FromTo(from int, to int) string {
	result := ""

	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "invalid \n"
	} else if from == to && from < 10 {
		result += "0" + strconv.Itoa(from) + "\n"
	}

	if from > 10 {
		for i := from; i >= 10; i++ {
			if i < 10 {
				result += "0"
			}
			result += strconv.Itoa(i)
			if i-1 >= 10 {
				result += ", "
			}

		}
		return result + "\n"
	}

	if from < 10 {
		for i := from; i <= 10; i++ {
			if i < 10 {
				result += "0"
			}
			result += strconv.Itoa(i)
			if i+1 >= 10 {
				result += ", "
			}

		}
	}
	return result + "\n"

}
