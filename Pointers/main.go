package main

import "fmt"

func main() {
	a := 10

	fmt.Println(a)
	//the below
	fmt.Printf("this is the address of a : %v", &a)
}
