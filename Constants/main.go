package main

import (
	"fmt"
	"math"
)

//in go u can make a string, boolean, and a numeric value a constant

const identity string = "learn2earn"

func main() {

	const k = 200

	a := k / 2

	result := float64(a)

	fmt.Println(identity, a, math.Sin(result))
}

//numbers can be int , uint , and float(decimal numbers),
//  u can perform mathematic arithmetic like sin, cos etc on float only
//u cant combine them during operations except that u combine them.
