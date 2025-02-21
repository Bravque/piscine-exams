// Write a program that takes strings as arguments, and displays its last
// argument followed by a newline.

// If the number of arguments is less than 1, the program displays a newline.

package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if there is at least one argument
	if len(os.Args) > 1 {
		// Print the last argument followed by a newline
		fmt.Println(os.Args[len(os.Args)-1])
	} else {
		// Print just a newline if no arguments are provided
		fmt.Println()
	}
}
