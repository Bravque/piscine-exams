// Write a program that takes a positive integer as argument and displays the sum
// of all prime numbers inferior or equal to it followed by a newline.

// If the number of arguments is not 1, or the argument is not a positive number,
// just display 0 followed by a newline.

// Yes, the examples are right.

package main

import (
	"fmt"
	"os"
)

// isPrime checks if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// Check divisibility from 2 to n/2 (since no number can be divisible by more than n/2 except itself)
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// stringToInt converts a string to an integer manually
func stringToInt(s string) (int, bool) {
	num := 0
	for i := 0; i < len(s); i++ {
		// Check if the character is a valid digit
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		num = num*10 + int(s[i]-'0') // Build the number digit by digit
	}
	return num, true
}

// main function to process the input and calculate the sum of primes
func main() {
	// Check if the number of arguments is 1
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	// Convert the argument to an integer
	num, valid := stringToInt(os.Args[1])
	if !valid || num <= 0 {
		// If the conversion fails or the number is not positive, print 0
		fmt.Println(0)
		return
	}

	// Sum all prime numbers <= num
	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	// Print the sum followed by a newline
	fmt.Println(sum)
}
