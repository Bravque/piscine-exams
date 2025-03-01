package main

import "fmt"

// NRune returns the nth rune in the string or 0 if it's out of bounds
func NRune(s string, n int) rune {
	// Check if n is a valid index
	if n < 0 || n >= len(s) {
		return 0
	}

	// Iterate over the string
	for i, char := range s {
		if i == n {
			return char
		}
	}

	// If n is out of bounds, return 0 (although the previous check prevents this case)
	return 0
}

func main() {
	// Test the NRune function and print the corresponding characters
	fmt.Println(string(NRune("Hello!", 3)))  // Output: 'l'
	fmt.Println(string(NRune("Salut!", 2)))  // Output: 'l'
	fmt.Println(string(NRune("Bye!", -1)))   // Output: (empty, as 0 is returned)
	fmt.Println(string(NRune("Bye!", 5)))    // Output: (empty, as 0 is returned)
	fmt.Println(string(NRune("Ola!", 4)))    // Output: (empty, as 0 is returned)
}
