// Reproduce the behavior of the function strdup (man strdup).

package main

import (
	"fmt"
)

func strdup(s string) string {
	// Return a copy of the input string
	return s
}

func main() {
	original := "Hello, Go!"
	duplicate := strdup(original)

	fmt.Println("Original:", original)
	fmt.Println("Duplicate:", duplicate)
}
