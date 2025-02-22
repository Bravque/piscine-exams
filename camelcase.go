package main

import (
	"fmt"
	"os"
)

func camelToSnake(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		c := s[i]

		// If the character is uppercase (A-Z), add an underscore before it (except at the start)
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				result += "_"
			}
			// Convert uppercase to lowercase
			c += 32
		}

		result += string(c)
	}

	return result
}

func main() {
	// Ensure an argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go camelCaseString")
		return
	}

	// Convert and print the snake_case result
	fmt.Println(camelToSnake(os.Args[1]))
}
