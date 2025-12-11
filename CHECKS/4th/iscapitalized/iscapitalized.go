package main
import "fmt"

/*

Write a function IsCapitalized() that takes a string as an argument and returns true if each word in the string begins with either an uppercase letter or a non-alphabetic character.

If any of the words begin with a lowercase letter return false.
If the string is empty return false.

*/

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}


func IsCapitalized(s string) bool {

	if s == ""{
		return false
	}

	for _, item:= range s {

		if item >= 'a' || item <= 'z'{
			
		}
	}

}