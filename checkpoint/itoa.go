package main

import "fmt"

func Itoa(n int) string {
	// Special case for 0
	if n == 0 {
		return "0"
	}

	// Handle negative numbers
	negative := false
	if n < 0 {
		negative = true
		n = -n // Work with absolute value
	}

	// Create a slice to collect digits in reverse order
	var result []byte

	// Extract digits and store them in reverse order
	for n > 0 {
		digit := n % 10
		result = append(result, byte(digit+'0')) // Convert digit to byte and append
		n = n / 10
	}

	// If the number was negative, prepend the '-' sign
	if negative {
		result = append(result, '-')
	}

	// Reverse the result slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// Convert the slice to string and return
	return string(result)
}

func main() {
    fmt.Println(Itoa(12345))       // "12345"
    fmt.Println(Itoa(0))           // "0"
    fmt.Println(Itoa(-1234))       // "-1234"
    fmt.Println(Itoa(987654321))  // "987654321"
}
