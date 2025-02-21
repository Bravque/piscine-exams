// Write a function that displays a string on the standard output.

// The pointer passed to the function contains the address of the string's first
// character.

package main

import "fmt"

// Function that displays a string using a pointer to its first character
func displayString(str *string) {
	// Dereference the pointer and print the string
	fmt.Println(*str)
}

func main() {
	// Create a string
	s := "Hello, Go!"

	// Pass a pointer to the string to the function
	displayString(&s)
}
