package main

import (
	"fmt"
	"os"
)

func alphamirror(s string) {
	// Iterate over each character in the string
	for _, ch := range s {
		// Check if the character is a lowercase letter
		if ch >= 'a' && ch <= 'z' {
			// Calculate the opposite character in lowercase
			ch = 'z' - (ch - 'a')
		} else if ch >= 'A' && ch <= 'Z' {
			// Calculate the opposite character in uppercase
			ch = 'Z' - (ch - 'A')
		}
		// Print the modified character
		fmt.Print(string(ch))
	}
	// Print a newline at the end
	fmt.Println()
}

func main() {
	// Check if there is exactly one argument
	if len(os.Args) != 2 {
		// If not, just print a newline
		fmt.Println()
	} else {
		// Otherwise, call the alphamirror function with the argument
		alphamirror(os.Args[1])
	}
}
