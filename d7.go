// Write a program that takes a positive (or zero) number expressed in base 10,
// and displays it in base 16 (lowercase letters) followed by a newline.

// If the number of parameters is not 1, the program displays a newline.


package main

import (
	"fmt"
	"os"
)

// Function to convert a number from base 10 to base 16
func toHex(num int) string {
	// If the number is 0, directly return "0"
	if num == 0 {
		return "0"
	}

	// Hexadecimal digits
	hexDigits := "0123456789abcdef"
	var result string

	// Repeatedly divide the number by 16 and store remainders
	for num > 0 {
		result = string(hexDigits[num%16]) + result
		num = num / 16
	}

	return result
}

func main() {
	// Check if the number of arguments is 1
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	// Convert the argument to an integer (base 10)
	var num int
	_, err := fmt.Sscanf(os.Args[1], "%d", &num)
	if err != nil || num < 0 {
		// If conversion fails or the number is negative, print a newline
		fmt.Println()
		return
	}

	// Convert the number to hexadecimal and print it
	fmt.Println(toHex(num))
}
