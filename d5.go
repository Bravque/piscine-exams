// Write a program named hidenp that takes two strings and displays 1
// followed by a newline if the first string is hidden in the second one,
// otherwise displays 0 followed by a newline.

// Let s1 and s2 be strings. We say that s1 is hidden in s2 if it's possible to
// find each character from s1 in s2, in the same order as they appear in s1.
// Also, the empty string is hidden in any string.

// If the number of parameters is not 2, the program displays a newline.


package main

import (
	"fmt"
	"os"
)

// Function to check if s1 is hidden in s2
func isHidden(s1, s2 string) bool {
	// Two pointers, one for each string
	i, j := 0, 0

	// Iterate through both strings
	for i < len(s1) && j < len(s2) {
		// If characters match, move the pointer for s1
		if s1[i] == s2[j] {
			i++
		}
		// Always move the pointer for s2
		j++
	}

	// If all characters of s1 are found in s2, return true
	return i == len(s1)
}

func main() {
	// Check if there are exactly two arguments
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	// Get the two strings from arguments
	s1 := os.Args[1]
	s2 := os.Args[2]

	// Check if s1 is hidden in s2
	if isHidden(s1, s2) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
