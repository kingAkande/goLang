package main 

import "fmt"

/**

Write a function called FishAndChips() that takes an int and returns a string.

If the number is divisible by 2, print fish.
If the number is divisible by 3, print chips.
If the number is divisible by 2 and 3, print fish and chips.
If the number is negative return error: number is negative.
If the number is non divisible by 2 or 3 return error: non divisible.

*/

func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}

func FishAndChips(arg int) string {
	if arg < 0 {
		return  "number is negative"
	}
	
	if arg%2==0 && arg%3 ==0 {
		return ("fish and chips")

	}
	if arg%2 == 0 {
		return ("fish")
	}
	if arg%3 == 0 {
		return ("chips")
	}


	
	 return  "non devisible"
}
