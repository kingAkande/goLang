package main

import "github.com/01-edu/go-tests/lib/is"

/**
Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.

If there are no primes inferior to the int passed as parameter the function should return 0.


**/

func FindPrevPrime(nb int) int {

	if nb < 2 {
		return 0
	}

	if is.Prime(nb) {
		return nb
	}
	return FindPrevPrime(nb - 1)
}
