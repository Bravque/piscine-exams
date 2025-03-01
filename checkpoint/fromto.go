package main

import "github.com/01-edu/z01"

func FromTo(from int, to int) {
	// Check if the 'from' or 'to' are out of the valid range
	if from < 0 || from > 99 || to < 0 || to > 99 {
		// Print "Invalid" if out of range
		z01.PrintRune('I')
		z01.PrintRune('n')
		z01.PrintRune('v')
		z01.PrintRune('a')
		z01.PrintRune('l')
		z01.PrintRune('i')
		z01.PrintRune('d')
		z01.PrintRune('\n')
		return
	}

	// Function to print an integer with leading zero if it's less than 10
	printNum := func(num int) {
		if num < 10 {
			z01.PrintRune('0')              // Print leading zero for numbers less than 10
		}
		z01.PrintRune(rune(num/10 + '0'))  // Print the tens place
		z01.PrintRune(rune(num%10 + '0'))  // Print the ones place
	}

	// Printing numbers from 'from' to 'to'
	if from < to {
		for i := from; i <= to; i++ {
			printNum(i)
			if i < to {
				z01.PrintRune(',')  // Print a comma
				z01.PrintRune(' ')  // Print a space
			}
		}
	} else if from > to {
		// Printing numbers from 'from' to 'to' in reverse order
		for i := from; i >= to; i-- {
			printNum(i)
			if i > to {
				z01.PrintRune(',')  // Print a comma
				z01.PrintRune(' ')  // Print a space
			}
		}
	} else {
		// If from == to, print the number
		printNum(from)
	}

	z01.PrintRune('\n')  // Print a newline at the end
}

func main() {
	FromTo(1, 10)  // 01, 02, 03, 04, 05, 06, 07, 08, 09, 10
	FromTo(10, 1)  // 10, 09, 08, 07, 06, 05, 04, 03, 02, 01
	FromTo(10, 10) // 10
	FromTo(100, 10) // Invalid
}
