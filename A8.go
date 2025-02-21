// Write a program that displays the alphabet in reverse, with even letters in
// uppercase, and odd letters in lowercase, followed by a newline.


package main

import "fmt"

func main() {
	// Loop through the alphabet in reverse order from 'z' to 'a'
	for i := 25; i >= 0; i-- {
		letter := 'a' + rune(i) // Calculate the letter at position i

		// If the index is even, print the letter in uppercase
		if i%2 == 0 {
			fmt.Print(string(letter - 'a' + 'A'))
		} else {
			// If the index is odd, print the letter in lowercase
			fmt.Print(string(letter))
		}
	}

	// Print a newline at the end
	fmt.Println()
}
