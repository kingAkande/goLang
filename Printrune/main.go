package main

import "github.com/01-edu/z01"

func main(){
		var aRune rune = 'D'

		alphababets :=[]string{"come" , "go "}
		digits := []rune{'d' , 'k'}
		var name string = "good boy"


	z01.PrintRune(aRune)
	z01.PrintRune(rune(name[0]))
	z01.PrintRune(rune(alphababets[1][1]))
	z01.PrintRune(digits[1])
}

