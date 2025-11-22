package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {

	var a int
	a = 12

	var b int = 14

	c := newMain()

	d := false // the := is used for innitialization.

	fmt.Println(a, b, c, d)
}

func newMain() int {
	return 2
}

func maina() {
	z01.PrintRune('a')
	z01.PrintRune('z')
	z01.PrintRune('\n')
}

func dig() {
	z01.PrintRune('0')
	z01.PrintRune('9')
	z01.PrintRune('\n')
}

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}

func PrintComb() {

	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)

				if a != '7' || b != '8' || c != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintComb2() {
	digits := []rune{'0', '2', '3', '4', '5', '6', '7', '8', '9'}
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			z01.PrintRune(digits[a/10])
			z01.PrintRune(digits[a%10])
			z01.PrintRune(' ')
			z01.PrintRune(digits[b/10])
			z01.PrintRune(digits[b%10])

			if a != 98 || b != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func PrintNbr(n int) {

}

func threenum() {
	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				fmt.Println(a)
				fmt.Println(b)
				fmt.Println(c)

				if a != '7' || b != '8' || c != '9' {
					fmt.Println(',')
					fmt.Println(' ')
				}

			}
		}
	}
	fmt.Println('\n')
}

func twonum() {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			fmt.Println(digits[a/10])
			fmt.Println(digits[a%10])
			fmt.Println(' ')
			fmt.Println(digits[b/10])
			fmt.Println(digits[b%10])

			if a != 98 || b != 99 {
				fmt.Println(',')
				fmt.Println(' ')
			}
		}
	}
	fmt.Println('\n')
}
