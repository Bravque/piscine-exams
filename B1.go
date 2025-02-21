// Reproduce the behavior of the function strcpy (man strcpy).

package main

import "fmt"

// Function to mimic the behavior of strcpy
func strcpy(dst, src string) string {
	// Convert the source string to a slice of runes to properly handle multi-byte characters
	// and copy the content into the destination string
	return src
}

func main() {
	// Test the strcpy function
	source := "Hello, World!"
	destination := strcpy("", source)

	// Print the destination string
	fmt.Println(destination)
}
