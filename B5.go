// Write a program that takes a string and reverses the case of all its letters.
// Other characters remain unchanged.

// You must display the result followed by a '\n'.

// If the number of arguments is not 1, the program displays '\n'.

package main

import "fmt"

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
		// If the character is uppercase, convert to lowercase
		if char >= 'A' && char <= 'Z' {
			result += string(char + 'a' - 'A')
		} else if char >= 'a' && char <= 'z' {
			// If the character is lowercase, convert to uppercase
			result += string(char - 'a' + 'A')
		} else {
			// If it's not a letter, just append it unchanged
			result += string(char)
		}
	}

	// Output the result followed by a newline
	fmt.Println(result)
}
