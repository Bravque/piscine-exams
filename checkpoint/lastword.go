package main

import "fmt"

func LastWord(s string) string {
	var result string
	ended := false

	// Start from the end of the string
	for i := len(s) - 1; i >= 0; i-- {
		word := s[i]

		// Skip trailing spaces
		if word == ' ' && !ended {
			continue
		}

		// Once we encounter a non-space character, start collecting the word
		if word == ' ' {
			break
		}

		// Build the word backwards
		result = string(word) + result
		ended = true
	}

	return result
}

func main() {
	fmt.Println(LastWord("hello there"))       
	fmt.Println(LastWord(""))                  
	fmt.Println(LastWord("hello ....... bye")) 
	fmt.Println(LastWord("    hello ....... bye    ")) 
}
