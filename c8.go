// Write a program called alpha_mirror that takes a string and displays this string
// after replacing each alphabetical character by the opposite alphabetical
// character, followed by a newline.

// 'a' becomes 'z', 'Z' becomes 'A'
// 'd' becomes 'w', 'M' becomes 'N'

// and so on.

// Case is not changed.

// If the number of arguments is not 1, display only a newline.

package main

import (
	"fmt"
	"os"
)

// alphaMirror takes a string and replaces each letter with its opposite
func alphaMirror(s string) string {
	var result []rune

	// Iterate over the string
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			// For lowercase letters, calculate the opposite
			result = append(result, 'a'+'z'-ch)
		} else if ch >= 'A' && ch <= 'Z' {
			// For uppercase letters, calculate the opposite
			result = append(result, 'A'+'Z'-ch)
		} else {
			// For non-alphabetical characters, just append as is
			result = append(result, ch)
		}
	}

	return string(result)
}

func main() {
	// Check if there is exactly one command-line argument
	if len(os.Args) != 2 {
		fmt.Println() // If not, print just a newline
		return
	}

	// Get the input string
	input := os.Args[1]

	// Apply the alpha_mirror transformation
	output := alphaMirror(input)

	// Print the result followed by a newline
	fmt.Println(output)
}
