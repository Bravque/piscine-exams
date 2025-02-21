
// Write a program that takes one or more strings and, for each argument, puts 
// the last character of each word (if it's a letter) in uppercase and the rest
// in lowercase, then displays the result followed by a \n.

// A word is a section of string delimited by spaces/tabs or the start/end of the
// string. If a word has a single letter, it must be capitalized.

// If there are no parameters, display \n.

package main

import (
	"fmt"
	"os"
)

// Function to convert a character to lowercase
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

// Function to convert a character to uppercase
func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - ('a' - 'A')
	}
	return c
}

// Function to process a word
func processWord(word []byte) []byte {
	// If the word has only one character, capitalize it
	if len(word) == 1 {
		word[0] = toUpper(word[0])
		return word
	}

	// Convert the entire word to lowercase
	for i := 0; i < len(word); i++ {
		word[i] = toLower(word[i])
	}

	// Capitalize the last character if it's a letter
	if word[len(word)-1] >= 'a' && word[len(word)-1] <= 'z' {
		word[len(word)-1] = toUpper(word[len(word)-1])
	}

	return word
}

func main() {
	// Check if there are no arguments
	if len(os.Args) == 1 {
		fmt.Println()
		return
	}

	// Iterate over each argument (starting from os.Args[1] as os.Args[0] is the program name)
	for i := 1; i < len(os.Args); i++ {
		// Convert the argument to a byte slice for manual processing
		arg := []byte(os.Args[i])
		var word []byte

		for j := 0; j < len(arg); j++ {
			// Check for spaces or tabs as word delimiters
			if arg[j] == ' ' || arg[j] == '\t'
