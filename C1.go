// Write a function that converts the string argument str to an integer (type int)
// and returns it.

// It works much like the standard atoi(const char *str) function, see the man.

package main

import (
	"fmt"
)

func ft_atoi(str string) int {
	// Initialize variables
	result := 0
	sign := 1
	i := 0

	// Skip any leading whitespaces
	for i < len(str) && (str[i] == ' ' || str[i] == '\t') {
		i++
	}

	// Check for a negative or positive sign
	if i < len(str) && str[i] == '-' {
		sign = -1
		i++
	} else if i < len(str) && str[i] == '+' {
		i++
	}

	// Convert the string to an integer
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		result = result*10 + int(str[i]-'0')
		i++
	}

	// Return the result with the correct sign
	return result * sign
}

func main() {
	// Example usage of the function
	str := "  -12345"
	num := ft_atoi(str)
	fmt.Println(num)  // Output: -12345

	str2 := "42"
	num2 := ft_atoi(str2)
	fmt.Println(num2) // Output: 42

	str3 := "  +987"
	num3 := ft_atoi(str3)
	fmt.Println(num3) // Output: 987
}
