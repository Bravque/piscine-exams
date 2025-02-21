// Write a program that takes a string and displays it, replacing each of its
// letters by the next one in alphabetical order.

// 'z' becomes 'a' and 'Z' becomes 'A'. Case remains unaffected.

// The output will be followed by a \n.

// If the number of arguments is not 1, the program displays \n.


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

	// Iterate over the string and shift each letter
	var result string
	for i := 0; i < len(str); i++ {
		// Get the current character
		c := str[i]

		// Handle lowercase letters (a-z)
		if c >= 'a' && c <= 'z' {
			if c == 'z' {
				result += "a" // Wrap around to 'a' if 'z'
			} else {
				result += string(c + 1) // Shift to the next letter
			}
		} 
		// Handle uppercase letters (A-Z)
		if c >= 'A' && c <= 'Z' {
			if c == 'Z' {
				result += "A" // Wrap around to 'A' if 'Z'
			} else {
				result += string(c + 1) // Shift to the next letter
			}
		}
		// If not a letter, just append it as is
	}

	// Print the result followed by a newline
	fmt.Println(result)
}
