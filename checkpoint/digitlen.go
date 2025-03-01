package main

import "fmt"

func DigitLen(n, base int) int {
	// Check if the base is valid (positive integer)
	if base <= 1 {
		return -1
	}

	// Use the absolute value of n to handle negative numbers
	if n < 0 {
		n = -n // Manually compute the absolute value
	}

	// Special case: if n is zero, the number of digits is 1 (e.g., "0")
	if n == 0 {
		return 1
	}

	// Calculate the number of digits by repeatedly dividing by the base
	digits := 0
	for n > 0 {
		n /= base
		digits++
	}

	return digits
}

func main() {
	// Test cases
	fmt.Println(DigitLen(100, 10))  // 3 digits in base 10
	fmt.Println(DigitLen(100, 2))   // 7 digits in base 2
	fmt.Println(DigitLen(-100, 16)) // 3 digits in base 16
	fmt.Println(DigitLen(100, -1))  // Error: Base must be greater than 1
}
