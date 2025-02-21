// Write a function that takes a byte, reverses it, bit by bit (like the
// 	example) and returns the result.
	

package main

import "fmt"

// reverseByte takes a byte and returns the byte with its bits reversed
func reverseByte(b byte) byte {
	var reversed byte

	// Iterate over each bit in the byte (8 bits)
	for i := 0; i < 8; i++ {
		// Shift the bits of the reversed byte to the left to make room for the new bit
		reversed <<= 1

		// Get the least significant bit of the original byte and add it to the reversed byte
		reversed |= b & 1

		// Shift the original byte to the right to process the next bit
		b >>= 1
	}

	return reversed
}

func main() {
	// Example byte (binary representation: 00000001)
	b := byte(1)

	// Reverse the bits of the byte
	reversed := reverseByte(b)

	// Print the result (binary representation of the reversed byte)
	fmt.Printf("Original byte: %08b\n", b)
	fmt.Printf("Reversed byte: %08b\n", reversed)
}
