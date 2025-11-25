package main

import "github.com/01-edu/z01"

/**
onlyf
Instructions
Write a program that displays an f character on the standard output and nothing else.

Usage
$ go run .
f
$ go run . "a" "b"
f
$ go run . "a" "b" "c"
f
*/

func main() {
	z01.PrintRune('f')
}
