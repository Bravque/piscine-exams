// Write a program that takes a string and displays it, replacing each of its
// letters by the letter 13 spaces ahead in alphabetical order.

// 'z' becomes 'm' and 'Z' becomes 'M'. Case remains unaffected.

// The output will be followed by a newline.

// If the number of arguments is not 1, the program displays a newline.

package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the number of arguments is exactly 1
	if len(os.Args) != 2 {
		fmt.Println() // Print a newline if not
		return
	}

	// Get the input string
	str := os.Args[1]

	// Create an empty result string
	var result string

	// Iterate through each character in the string
	for _, char := range str {
		// Check if the character is a lowercase letter
		if char >= 'a' && char <= 'z' {
			// Rotate the character by 13 positions within 'a' to 'z'
			result += string((char-'a'+13)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			// Rotate the character by 13 positions within 'A' to 'Z'
			result += string((char-'A'+13)%26 + 'A')
		} else {
			// If it's not a letter, just append it unchanged
			result += string(char)
		}
	}

	// Output the result followed by a newline
	fmt.Println(result)
}
