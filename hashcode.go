package main

import (
	"fmt"
	"os"
)

func hashCode(s string) int {
	hash := 0
	for i := 0; i < len(s); i++ {
		hash = hash*31 + int(s[i]) // Multiply by 31 (common hashing prime) and add char value
	}
	return hash
}

func main() {
	// Ensure an argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go yourString")
		return
	}

	// Generate and print the hash code
	fmt.Println(hashCode(os.Args[1]))
}
