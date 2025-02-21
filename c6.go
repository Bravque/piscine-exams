// Write a function that takes a byte, swaps its halves (like the example) and
// returns the result.

package main

import "fmt"

// swapByteHalves takes a byte, swaps its upper and lower 4 bits, and returns the result
func swapByteHalves(b byte) byte {
	// Extract the upper 4 bits by shifting right and the lower 4 bits by masking with 0x0F
	upper := b >> 4          // Get the upper 4 bits
	lower := b & 0x0F        // Get the lower 4 bits

	// Swap the halves by shifting the lower bits to the higher 4 positions
	// and the upper bits to the lower 4 positions
	return (lower << 4) | upper
}

func main() {
	// Example byte (binary: 11010101)
	b := byte(213)

	// Swap the halves of the byte
	swapped := swapByteHalves(b)

	// Print the result (binary representation of the original and swapped byte)
	fmt.Printf("Original byte: %08b\n", b)
	fmt.Printf("Swapped byte:  %08b\n", swapped)
}
