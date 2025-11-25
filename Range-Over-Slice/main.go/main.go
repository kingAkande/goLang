package main

import "fmt"

func main() {

	name := []string{"sualimon", "jibril", "muhammad"}

	for index, _ := range name {

		fmt.Print(index)
	}

	for _, item := range name {
		fmt.Println(item)
	}

	// fmt.Println(name[0])

	gdp(48, 18)
}

func gdp(x, y int) {

	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for index, item := range num {
		// fmt.Println(index , item)

		if x%num[index] == 0 && y%num[index] == 0 {

			fmt.Println(item)

		}
	}

}
