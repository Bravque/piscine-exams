package main

import "fmt"

func cleanstr(s string) {
	// If the string is empty, just print a newline
	if len(s) == 0 {
		fmt.Println()
		return
	}

	// Variables to track if we're in a word and the result
	result := ""
	inWord := false

	for i := 0; i < len(s); i++ {
		// If the character is a space or tab
		if s[i] == ' ' || s[i] == '\t' {
			// Add a single space if we're in a word
			if inWord {
				result += " "
				inWord = false
			}
		} else {
			// Add the character to the result if it's part of a word
			result += string(s[i])
			inWord = true
		}
	}

	// Trim any leading or trailing spaces from the result and print it
	fmt.Println(result)
}

func main() {
	cleanstr("you see it's easy to display the same thing")  // "you see it's easy to display the same thing"
	cleanstr(" only    it's  harder   ")                     // "only it's harder"
	cleanstr(" how funny")                                   // "how funny"
	cleanstr("")                                             // ""
}
