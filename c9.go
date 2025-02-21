
// Write the following function:

// int		max(int* tab, unsigned int len);

// The first parameter is an array of int, the second is the number of elements in
// the array.

// The function returns the largest number found in the array.

// If the array is empty, the function returns 0.

package main

import "fmt"

// max takes a slice of integers and returns the largest value
func max(tab []int) int {
	// If the slice is empty, return 0
	if len(tab) == 0 {
		return 0
	}

	// Initialize max_val with the first element of the slice
	maxVal := tab[0]

	// Loop through the slice to find the maximum value
	for _, value := range tab[1:] {
		if value > maxVal {
			maxVal = value
		}
	}

	return maxVal
}

func main() {
	// Example slice
	arr := []int{1, 2, 3, 4, 5}

	// Call the max function and print the result
	result := max(arr)
	fmt.Printf("The maximum value is: %d\n", result)

	// Test with an empty slice
	emptyArr := []int{}
	resultEmpty := max(emptyArr)
	fmt.Printf("The maximum value in the empty array is: %d\n", resultEmpty)
}
