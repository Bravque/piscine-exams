// Write a program that takes a string and displays its first word, followed by a
// newline.

// A word is a section of string delimited by spaces/tabs or by the start/end of
// the string.

// If the number of parameters is not 1, or if there are no words, simply display
// a newline.

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

	// Find the first word by checking for spaces or tabs
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || str[i] == '\t' { // Found a space or tab
			break
		}
		word += string(str[i]) // Add the character to the word
	}

	// If word is empty, print a newline
	if word == "" {
		fmt.Println()
		return
	}

	// Print the first word followed by a newline
	fmt.Println(word)
}
