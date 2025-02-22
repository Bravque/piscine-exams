// Write a program called search_and_replace that takes 3 arguments, the first 
// arguments is a string in which to replace a letter (2nd argument) by
// another one (3rd argument).

// If the number of arguments is not 3, just display a newline.

// If the second argument is not contained in the first one (the string)
// then the program simply rewrites the string followed by a newline.

package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the number of arguments is exactly 3
	if len(os.Args) != 4 {
		fmt.Println() // Print a newline if not
		return
	}

	// Convert arguments to runes
	input := []rune(os.Args[1])
	oldChar := []rune(os.Args[2])
	newChar := []rune(os.Args[3])

	// Ensure the oldChar and newChar are single characters
	if len(oldChar) != 1 || len(newChar) != 1 {
		fmt.Println()
		return
	}

	// Replace occurrences of oldChar with newChar
	for i, char := range input {
		if char == oldChar[0] {
			input[i] = newChar[0]
		}
	}

	// Output the result followed by a newline
	fmt.Println(string(input))
}
