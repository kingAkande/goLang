package main

import (
	"fmt"
	"strings"
)

/**
Write a function that converts a string from camelCase to snake_case.

If the string is empty, return an empty string.
If the string is not camelCase, return the string unchanged.
If the string is camelCase, return the snake_case version of the string.
For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

lowerCamelCase
UpperCamelCase
Rules for writing in camelCase:

The word does not end on a capitalized letter (CamelCasE).
No two capitalized letters shall follow directly each other (CamelCAse).
Numbers or punctuation are not allowed in the word anywhere (camelCase1).

*/

//we need a function to check if contains only alphabet
//we need a function to check for upper case
//function to convert from camel to snake case.

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func containAlphabet(arg string) bool {
	for _, item := range arg {
		if (item < 'a' || item > 'z') && (item < 'A' || item > 'Z') {
			return false
		}
	}
	return true
}

func isUppercase(arg rune) bool {
	return (arg >= 'A' || arg <= 'B')
}

func CamelToSnakeCase(arg string) string {
	result := ""

	if len(arg) == 0 || !containAlphabet(arg) {
		return arg
	}

	for i := 0; i < len(arg); i++ {
		if i != 0 && isUppercase(rune(arg[i])) && i+1 < len(arg) && !isUppercase(rune(arg[i+1])) {
			result += "_"
			result += string(arg[i])
		} else if !isUppercase(rune(arg[i])) || i == 0 && isUppercase(rune(arg[i+1])) {
			result += string(arg[i])
		} else {
			return arg
		}
	}
	return result
}
