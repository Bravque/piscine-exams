package main

import "fmt"

func FirstRune(s string) rune {
	// Iterate over the string from the beginning
	for _, char := range s {
		if char != ' ' { // Skip spaces
			return char
		}
	}
	// If no non-space characters are found, return a zero value rune
	return 0
}

func main() {
	fmt.Println(string(FirstRune("Hello!")))   // Output: H
	fmt.Println(string(FirstRune("   Salut!"))) // Output: S
	fmt.Println(string(FirstRune("   Ola!   "))) // Output: O
	fmt.Println(string(FirstRune("   ")))       // Output: (no output, returns 0)
}
