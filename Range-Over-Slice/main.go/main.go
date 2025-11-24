package main

import "fmt"

func main() {

	name := []string{"sualimon", "jibril", "muhammad"}

	for index, _ := range name {

		fmt.Println(index)
	}

	for _, item := range name {
		println(item)
	}

	// fmt.Println(name[0])

}
