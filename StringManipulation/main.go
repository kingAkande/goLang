package main

// A string is a sequence of bytes

import (
	"fmt"
	"strings"
)

func main(){
	name:= "ILYAS"

	speach := "this is you"

	first := name[0]
	convertName := []byte(name)

	convertName[0] = 'i'

	alteredName := string(convertName)

	fmt.Println(name)
	fmt.Println(convertName)
	fmt.Println(alteredName)
	fmt.Println(len(name))
	fmt.Println(strings.Split(speach, " ")[2])
	fmt.Println(string(first))

	loopString("come here")

	
}


func loopString(arg string)  {
	for _ , item:= range arg{
		fmt.Println(string(item))
	}
}