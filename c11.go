// Write a program that takes three strings:
// - The first and the third one are representations of base-10 signed integers
//   that fit in an int.
// - The second one is an arithmetic operator chosen from: + - * / %

// The program must display the result of the requested arithmetic operation,
// followed by a newline. If the number of parameters is not 3, the program
// just displays a newline.

// You can assume the string have no mistakes or extraneous characters. Negative
// numbers, in input or output, will have one and only one leading '-'. The
// result of the operation fits in an int.

package main

import (
	"fmt"
	"os"
)

// helper function to convert string to integer manually
func stringToInt(s string) (int, bool) {
	// Handle empty string case
	if len(s) == 0 {
		return 0, false
	}

	// Handle negative numbers
	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:] // Remove the negative sign for easier conversion
	}

	// Now convert the rest of the string
	var result int
	for i := 0; i < len(s); i++ {
		// Ensure the character is a digit between '0' and '9'
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		result = result*10 + int(s[i]-'0') // Build the number
	}

	if negative {
		result = -result
	}

	return result, true
}

// performOperation performs the arithmetic operation based on the operator
func performOperation(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		// Handle division by zero
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			os.Exit(1)
		}
		return num1 / num2
	case "%":
		// Handle modulo by zero
		if num2 == 0 {
			fmt.Println("Error: Modulo by zero")
			os.Exit(1)
		}
		return num1 % num2
	default:
		// Invalid operator
		fmt.Println("Invalid operator")
		os.Exit(1)
		return 0
	}
}

func main() {
	// Check if the number of arguments is exactly 3
	if len(os.Args) != 4 {
		fmt.Println() // If not, print a newline and return
		return
	}

	// Get the input arguments
	strNum1 := os.Args[1]
	operator := os.Args[2]
	strNum2 := os.Args[3]

	// Convert the first and third strings to integers using our custom function
	num1, valid1 := stringToInt(strNum1)
	num2, valid2 := stringToInt(strNum2)

	if !valid1 || !valid2 {
		fmt.Println() // If there is an error in conversion, print a newline
		return
	}

	// Perform the requested arithmetic operation and print the result
	result := performOperation(num1, num2, operator)
	fmt.Println(result)
}
