// Write a program that takes a string and displays its last word followed by a \n.

// A word is a section of string delimited by spaces/tabs or by the start/end of
// the string.

// If the number of parameters is not 1, or there are no words, display a newline.

package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if there is exactly one command-line argument
	if len(os.Args) != 2 {
		fmt.Println() // If not, print just a newline
		return
	}

	// Get the input string
	input := os.Args[1]

	// Initialize variables to keep track of the last word
	var lastWord string
	inWord := false

	// Iterate over the string from the end to find the last word
	for i := len(input) - 1; i >= 0; i-- {
		ch := input[i]

		// Check if the current character is a space or tab
		if ch == ' ' || ch == '\t' {
			if inWord {
				break // Exit once we finish the last word
			}
		} else {
			inWord = true
			lastWord = string(ch) + lastWord // Prepend character to build the word
		}
	}

	// If no word was found, print just a newline
	if lastWord == "" {
		fmt.Println()
	} else {
		// Print the last word followed by a newline
		fmt.Println(lastWord)
	}
}
