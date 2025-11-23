package main

import "fmt"

func main(){
	aString := "hello"

	manipulateString := []byte(aString)

	 manipulateString[0] = 'A'

	alterAletterIntheString := string(manipulateString)


	fmt.Println(aString)

	fmt.Println(manipulateString)
	fmt.Println(alterAletterIntheString)

}

