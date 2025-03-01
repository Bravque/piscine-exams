package main

import (
	"fmt"
)

func IsCapitalized(s string) bool {
	// If the string is empty, return false
	if len(s) == 0 {
		return false
	}

	// Variable to keep track of whether we are inside a word
	inWord := false

	// Iterate through each character of the string
	for i := 0; i < len(s); i++ {
		c := s[i]

		// Check if the current character is a space
		if c == ' ' {
			// If the previous character was a non-space, we were at the end of a word
			inWord = false
			continue
		}

		// If we're at the beginning of a word, check the first character
		if !inWord {
			// Mark that we've started a word
			inWord = true

			// Check if the first character is either uppercase or a non-alphabetic character
			if !(c >= 'A' && c <= 'Z' || c < 'A' || c > 'z') {
				// If the first character is lowercase, return false
				return false
			}
		}
	}

	// If all words pass the checks, return true
	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?")) // false
	fmt.Println(IsCapitalized("Hello How Are You"))  // true
	fmt.Println(IsCapitalized("Whats 4this 100K?"))  // true
	fmt.Println(IsCapitalized("Whatsthis4"))          // true
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))      // true
	fmt.Println(IsCapitalized(""))                    // false
}
