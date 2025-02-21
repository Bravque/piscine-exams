// Write a function that takes a byte, and prints it in binary WITHOUT A NEWLINE
// AT THE END.

package main

import "fmt"

// printByteInBinary prints the binary representation of a byte without a newline
func printByteInBinary(b byte) {
	for i := 7; i >= 0; i-- {
		// Print the bit at position i
		if (b>>i)&1 == 1 {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}

func main() {
	// Example usage
	var b byte = 5  // Example byte (in binary: 00000101)
	printByteInBinary(b)
}
