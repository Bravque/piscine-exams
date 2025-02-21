package main

import "fmt"

func main() {
	// Loop from 9 to 0 and print each digit
	for i := 9; i >= 0; i-- {
		fmt.Print(i)
	}
	// Print a newline after all digits
	fmt.Println()
}

// Write a program that displays all digits in descending order, followed by a
// newline.