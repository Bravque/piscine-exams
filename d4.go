// Write the following function:

// int     *ft_rrange(int start, int end);

// It must allocate (with malloc()) an array of integers, fill it with consecutive
// values that begin at end and end at start (Including start and end !), then
// return a pointer to the first value of the array.

package main

import "fmt"

// ftRrange function creates a slice of integers starting from `end` and
// ending at `start` (inclusive), and returns the slice.
func ftRrange(start, end int) []int {
	// Calculate the length of the slice
	length := end - start + 1
	
	// Create a slice to hold the range of integers
	result := make([]int, length)

	// Fill the slice with values from `end` to `start`
	for i := 0; i < length; i++ {
		result[i] = end - i
	}

	// Return the slice
	return result
}

func main() {
	// Example usage
	start := 3
	end := 7

	result := ftRrange(start, end)
	fmt.Println(result) // Output: [7 6 5 4 3]
}
