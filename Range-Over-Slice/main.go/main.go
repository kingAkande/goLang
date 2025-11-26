package main

import "fmt"

// import "fmt"

func main() {

	name := []string{"sualimon", "jibril", "muhammad"}

	for index, _ := range name {

		fmt.Print(index)
	}

	for _, item := range name {
		fmt.Println(item)
	}

	fmt.Println(name[0])

	// fmt.Println(gdp(100, 50))

	fmt.Println(Gcd(100, 50))
}

func gdp(x, y int) int {

	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := 1

	for index, item := range num {
		// fmt.Println(index , item)

		if x%num[index] == 0 && y%num[index] == 0 {

			// fmt.Println(item)
			result = item

		}
		}

	return result

}

func Gcd(a, b uint) uint {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

