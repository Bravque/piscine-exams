// Write a program that takes a string and displays it with exactly three spaces
// between each word, with no spaces or tabs either at the beginning or the end,
// followed by a newline.

// A word is a section of string delimited either by spaces/tabs, or by the
// start/end of the string.

// If the number of parameters is not 1, or if there are no words, simply display
// a newline.

package main

import (
	"fmt"
	"os"
)

// Function to split the string by spaces/tabs manually and return the words
func splitWords(input string) []string {
	var words []string
	var currentWord []byte

	for i := 0; i < len(input); i++ {
		// If the character is a space or tab, treat it as a delimiter
		if input[i] == ' ' || input[i] == '\t' {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = nil
			}
		} else {
			// Add character to the current word
			currentWord = append(currentWord, input[i])
		}
	}

	// Add the last word if there was one
	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	return words
}

func main() {
	// Check if the number of arguments is exactly 1
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	// Get the input string
	input := os.Args[1]

	// Split the input into words
	words := splitWords(input)

	// If no words were found, print a newline
	if len(words) == 0 {
		fmt.Println()
		return
	}

	// Print the words with exactly three spaces between each word
	for i, word := range words {
		if i > 0 {
			// Print three spaces between words
			fmt.Print("   ")
		}
		fmt.Print(word)
	}

	// Print a newline at the end
	fmt.Println()
}
