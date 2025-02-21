package main

import "fmt"

// reverseString reverses a string in-place and returns it
func reverseString(s string) string {
	// Convert the string to a slice of runes to handle multi-byte characters
	runes := []rune(s)
	
	// Reverse the slice of runes in-place
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	
	// Convert the slice of runes back to a string and return it
	return string(runes)
}

func main() {
	// Test the reverseString function
	str := "Hello, World!"
	reversedStr := reverseString(str)
	fmt.Println("Original:", str)
	fmt.Println("Reversed:", reversedStr)
}
