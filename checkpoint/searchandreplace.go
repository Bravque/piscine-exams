package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if there are exactly 3 arguments
	if len(os.Args) != 4 {
		return
	}

	str := os.Args[1]
	oldchar := os.Args[2]
	newchar := os.Args[3]

	if len(oldchar) != 1 || len(newchar) != 1 {
		return
	}

	// Convert the string to a slice of runes
	runes := []rune(str)

	// Convert oldchar and newchar to runes
	oldRune := []rune(oldchar)[0]
	newRune := []rune(newchar)[0]

	// Iterate over each rune and replace oldRune with newRune
	for i := 0; i < len(runes); i++ {
		if runes[i] == oldRune {
			runes[i] = newRune
		}
	}

	// Convert the rune slice back to a string and print the result
	fmt.Println(string(runes))
}
