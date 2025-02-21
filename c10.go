// Write a program that takes two strings and checks whether it's possible to
// write the first string with characters from the second string, while respecting
// the order in which these characters appear in the second string.

// If it's possible, the program displays the string, followed by a \n, otherwise
// it simply displays a \n.

// If the number of arguments is not 2, the program displays a \n.

package main

import (
	"fmt"
	"os"
)

// canFormString checks if it's possible to form `str1` using characters from `str2`
// while respecting the order of characters in `str2`.
func canFormString(str1, str2 string) bool {
	// Index for the current character in str1
	i := 0
	// Iterate through str2 and try to match str1 in order
	for _, ch := range str2 {
		if i < len(str1) && ch == rune(str1[i]) {
			// If the characters match, move to the next character in str1
			i++
		}
		// If all characters in str1 are found, break early
		if i == len(str1) {
			return true
		}
	}
	// If we've gone through str2 and not matched all characters in str1, return false
	return i == len(str1)
}

func main() {
	// Check if there are exactly two arguments
	if len(os.Args) != 3 {
		fmt.Println() // If not, print a newline and return
		return
	}

	// Get the two input strings
	str1 := os.Args[1]
	str2 := os.Args[2]

	// Check if it's possible to form str1 from str2
	if canFormString(str1, str2) {
		fmt.Println(str1) // If possible, display str1
	} else {
		fmt.Println() // Otherwise, just print a newline
	}
}

