package main

import (
	"fmt"
	"hash"
)

/**

Write a function called HashCode() that takes a string as an argument and returns a new hashed string.

The hash equation is computed as follows:
(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

If the resulting character is unprintable add 33 to it.

***/

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(arg string) string {

	size := len(arg)

	hashed := ""

	for _, item := range arg {
		hash := (int(item) + size) % 127

		if hash < 32 || hash > 126 {
			hash += 33
		}
		hashed += string(hash)
	}

	return hashed

}

