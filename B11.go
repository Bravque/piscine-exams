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

	// Iterate through the string to find the first word
	var word string
	for i := 0; i < len(str); i++ {
		// Check if the character is a space or tab (delimiter)
		if str[i] == ' ' || str[i] == '\t' {
			break
		}
		// Append the character to the word
		word += string(str[i])
	}

	// If a word is found, print it, else print a newline
	if len(word) > 0 {
		fmt.Println(word)
	} else {
		fmt.Println()
	}
}
