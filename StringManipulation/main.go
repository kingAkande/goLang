package main

import "fmt"

func main(){
	name:= "ILYAS"

	convertName := []byte(name)

	convertName[0] = 'i'

	alteredName := string(convertName)

	fmt.Println(name)
	fmt.Println(convertName)
	fmt.Println(alteredName)
}

