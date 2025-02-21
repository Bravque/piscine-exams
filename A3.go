package main

import "fmt"

// ftPrintNumbers prints digits from 0 to 9 in ascending order
func ftPrintNumbers() {
	for i := 0; i <= 9; i++ {
		fmt.Print(i)
	}
	fmt.Println() // Print a newline after the digits
}

func main() {
	// Call the function to print numbers
	ftPrintNumbers()
}

// Write a function that displays all digits in ascending order.
