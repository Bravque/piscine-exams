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
	"strings"
)

func main() {
	// Check if the number of arguments is exactly 3
	if len(os.Args) != 4 {
		fmt.Println() // Print a newline if not
		return
	}

	// Get the string, old character, and new character
	str := os.Args[1]
	oldChar := os.Args[2]
	newChar := os.Args[3]

	// Replace the old character with the new character if found
	if strings.Contains(str, oldChar) {
		// Replace all occurrences of oldChar with newChar
		str = strings.ReplaceAll(str, oldChar, newChar)
	}

	// Output the result followed by a newline
	fmt.Println(str)
}
