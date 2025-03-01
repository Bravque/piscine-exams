package main

import "fmt"
import "os"

func main() {
	// Ensure exactly one argument
	if len(os.Args) != 2 {
		return
	}

	// Get the input string
	input := os.Args[1]

	// Initialize variables
	var result, word string

	// Iterate over the input string to process each character
	for i := 0; i < len(input); i++ {
		c := input[i]

		if c == ' ' {
			// When space is encountered, add the word to the result (if it's not empty)
			if word != "" {
				if result != "" {
					result += "   " // Add exactly three spaces between words
				}
				result += word
				word = "" // Reset word
			}
		} else {
			// Add characters to the current word
			word += string(c)
		}
	}
	// Add the last word (if any) to the result
	if word != "" {
		if result != "" {
			result += "   "
		}
		result += word
	}

	// Print the result
	if result != "" {
		fmt.Println(result)
	}
}
