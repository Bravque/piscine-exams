// Write a program that takes a string, and displays the string in reverse
// followed by a newline.

// If the number of parameters is not 1, the program displays a newline.

package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if the number of arguments is exactly 1
	if len(os.Args) != 2 {
		fmt.Println() // Print a newline if the number of arguments is not 1
		return
	}

	// Get the input string
	str := os.Args[1]

	// Reverse the string
	var reversed string
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	// Print the reversed string followed by a newline
	fmt.Println(reversed)
}
