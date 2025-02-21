// Write a program that takes two strings representing two strictly positive
// integers that fit in an int.

// Display their highest common denominator followed by a newline (It's always a
// strictly positive integer).

// If the number of parameters is not 2, display a newline.

package main

import (
	"fmt"
	"os"
)

// Function to calculate the GCD using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to convert string to int without using strconv
func stringToInt(str string) int {
	result := 0
	for i := 0; i < len(str); i++ {
		result = result*10 + int(str[i]-'0')
	}
	return result
}

func main() {
	// Check if there are exactly two arguments
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	// Get the two input strings
	str1 := os.Args[1]
	str2 := os.Args[2]

	// Convert strings to integers
	num1 := stringToInt(str1)
	num2 := stringToInt(str2)

	// Calculate the GCD of num1 and num2
	result := gcd(num1, num2)

	// Print the result followed by a newline
	fmt.Println(result)
}
