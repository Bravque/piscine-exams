package main

import "fmt"

// Helper function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Function to find the largest prime number less than or equal to `nb`
func FindPrevPrime(nb int) int {
	// Start from nb and check each number downwards
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	// If no prime number is found, return 0
	return 0
}

func main() {
	// Test the function
	fmt.Println(FindPrevPrime(5))  // Expected output: 5
	fmt.Println(FindPrevPrime(4))  // Expected output: 3
}
