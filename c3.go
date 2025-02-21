// Write a program that takes two strings and displays, without doubles, the
// characters that appear in both strings, in the order they appear in the first
// one.

// The display will be followed by a \n.

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

	// Create a map to track characters that have already been printed
	seen := make(map[rune]bool)

	// Iterate through each character in the first string
	for _, ch1 := range str1 {
		// Iterate through each character in the second string
		for _, ch2 := range str2 {
			// If the character from str1 matches one from str2 and hasn't been printed
			if ch1 == ch2 && !seen[ch1] {
				// Print the character and mark it as seen
				fmt.Print(string(ch1))
				seen[ch1] = true
				break // Move to the next character in str1 once we've found a match
			}
		}
	}

	// Print a newline after the output
	fmt.Println()
}
