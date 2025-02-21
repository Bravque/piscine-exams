// Write a function that swaps the contents of two integers the adresses of which
// are passed as parameters.

package main

import "fmt"

// Function to swap the values of two integers
func swap(a *int, b *int) {
	// Swap the values pointed to by a and b
	*a, *b = *b, *a
}

func main() {
	// Declare two integers
	x := 10
	y := 20

	// Print values before swap
	fmt.Println("Before swap:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)

	// Call the swap function
