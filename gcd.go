package main

import (
	"fmt"
	"os"
)

// Function to compute GCD using the Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var num1, num2 int

	// Ensure two arguments are provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go num1 num2")
		return
	}

	// Read user input as integers
	_, err := fmt.Sscanf(os.Args[1], "%d", &num1)
	_, err2 := fmt.Sscanf(os.Args[2], "%d", &num2)

	// Check for conversion errors
	if err != nil || err2 != nil {
		fmt.Println("Please enter valid integers")
		return
	}

	// Calculate and print the GCD
	fmt.Println(gcd(num1, num2))
}
