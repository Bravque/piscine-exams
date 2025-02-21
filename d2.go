// Write a program that takes a string, and displays this string with exactly one
// space between words, with no spaces or tabs either at the beginning or the end,
// followed by a \n.

// A "word" is defined as a part of a string delimited either by spaces/tabs, or
// by the start/end of the string.

// If the number of arguments is not 1, or if there are no words to display, the
// program displays \n.
package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the number of arguments is 1
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	// Get the input string
	input := os.Args[1]
	var result []byte
	inWord := false

	// Iterate over the input string
	for i := 0; i < len(input); i++ {
		char := input[i]

		// Skip spaces and tabs at the beginning
		if char == ' ' || char == '\t' {
			if inWord {
				// Add a space if we are in a word (to separate words)
				result = append(result, ' ')
				inWord = false
			}
		} else {
			// Add the character to the result
			result = append(result, char)
			inWord = true
		}
	}

	// Handle case where there were words and result is built
	// Remove trailing space, if any
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	// Print the result followed by a newline
	fmt.Println(string(result))
}
