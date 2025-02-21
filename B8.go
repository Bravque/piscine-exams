// Write a program that prints the numbers from 1 to 100, each separated by a
// newline.

// If the number is a multiple of 3, it prints 'fizz' instead.

// If the number is a multiple of 5, it prints 'buzz' instead.

// If the number is both a multiple of 3 and a multiple of 5, it prints 'fizzbuzz' instead.

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		// Check if the number is a multiple of both 3 and 5
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			// Check if the number is a multiple of 3
			fmt.Println("fizz")
		} else if i%5 == 0 {
			// Check if the number is a multiple of 5
			fmt.Println("buzz")
		} else {
			// If it's neither, just print the number
			fmt.Println(i)
		}
	}
}
