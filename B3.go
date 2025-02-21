// Write a program called repeat_alpha that takes a string and display it
// repeating each alphabetical character as many times as its alphabetical index,
// followed by a newline.

package main

import "fmt"

func repeatAlpha(s string) {
	// Loop through each character in the string
	for _, char := range s {
		// Check if the character is a lowercase letter ('a' to 'z')
		if char >= 'a' && char <= 'z' {
			// Find the alphabetical index (1-based)
			index := int(char - 'a' + 1)
			
			// Print the character `index` times
			for i := 0; i < index; i++ {
				fmt.Print(string(char))
			}
		}
		// Check if the character is an uppercase letter ('A' to 'Z')
		if char >= 'A' && char <= 'Z' {
			// Convert to lowercase for index calculation
			index := int(char - 'A' + 1)

			// Print the character `index` times
			for i := 0; i < index; i++ {
				fmt.Print(string(char))
			}
		}
	}
	
	// Print a newline after the output
	fmt.Println()
}

func main() {
	// Test the repeatAlpha function
	inputString := "abcZ"
	repeatAlpha(inputString)
}
