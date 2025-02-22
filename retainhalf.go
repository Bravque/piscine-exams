package main

import (
	"fmt"
	"os"
)

func firstHalf(s string) string {
	// halfLength := len(s) / 2
	return s[:(len(s)/2)] // Slice to retain only the first half
}

func main() {
	// Ensure an argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go yourString")
		return
	}

	// Get the input and print the first half
	fmt.Println(firstHalf(os.Args[1]))
}
