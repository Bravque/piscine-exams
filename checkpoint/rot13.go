package main

import (
	"fmt"
	"os"
)

func main() {
	// Ensure there is at least one argument (excluding the program name)
	if len(os.Args) < 2 {
		return
	}

	// Get the input string
	args := os.Args[1]

	// Initialize the result as an empty string
	result := ""

	for _, char := range args {
		if char >= 'a' && char <= 'z' {
			// Apply ROT13 transformation for lowercase letters
			char = (char-'a'+13)%26 + 'a'
		} else if char >= 'A' && char <= 'Z' {
			// Apply ROT13 transformation for uppercase letters
			char = (char-'A'+13)%26 + 'A'
		}

		// Append the transformed character to the result
		result += string(char)
	}

	// Print the result
	fmt.Println(result)
}
