// Write a program that takes two strings and displays, without doubles, the
// characters that appear in either one of the strings.

// The display will be in the order characters appear in the command line, and
// will be followed by a \n.

// If the number of arguments is not 2, the program displays \n.

package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the number of arguments is exactly 2
	if len(os.Args) != 3 {
		fmt.Println() // If not, print just a newline
		return
	}

	// Get the two input strings
	str1 := os.Args[1]
	str2 := os.Args[2]

	// Create a map to track characters that have already been displayed
	seen := make(map[rune]bool)

	// Iterate over both strings and print characters without duplicates
	for _, ch := range str1 + str2 {
		// If the character hasn't been printed before
		if !seen[ch] {
			// Print the character and mark it as seen
			fmt.Print(string(ch))
			seen[ch] = true
		}
	}

	// Print a newline after the output
	fmt.Println()
}
