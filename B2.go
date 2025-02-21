// Write a function that returns the length of a string.

package main

import "fmt"

// Function to return the length of a string
func myStrLen(s string) int {
	// Iterate through the string and count characters
	length := 0
	for range s {
		length++
	}
	return length
}

func main() {
	// Test the myStrLen function
	testString := "Hello, World!"
	fmt.Println("Length of string:", myStrLen(testString))
}
