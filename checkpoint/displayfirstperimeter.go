package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	str := os.Args[1:] // Get arguments passed to the program

	// Check if there's at least one argument
	if len(str) > 0 {
		// Print the first rune (character) of the first argument
		for _, r := range str[0] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	} else {
		// If no arguments, print a default message (in rune form)
		for _, r := range "No arguments provided." {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
