package main

import "fmt"

func main() {
	i := 0
	for i <= 7 {
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}

	for i := range 5 {
		fmt.Println(i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for w := range 10 {
		if w%2 == 0 {
			fmt.Println(w)
		}

	}

}
