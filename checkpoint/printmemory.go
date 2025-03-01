package main

import "fmt"

func PrintMemory(arr [10]byte) {
	// Print the hexadecimal representation
	for i := 0; i < 10; i++ {
		// Print each byte in hexadecimal format
		fmt.Printf("%02x ", arr[i])
		if i == 7 {
			fmt.Println() // After printing the 8th byte, move to the next line
		}
	}

	// Print the ASCII representation
	for i := 0; i < 10; i++ {
		// Check if the character is printable
		if arr[i] >= 32 && arr[i] <= 126 {
			fmt.Printf("%c", arr[i])
		} else {
			fmt.Print(".") // Replace non-printable characters with a dot
		}
	}
	fmt.Println() // Print a newline at the end
}

func main() {
	// Test the PrintMemory function with example input
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*', 0, 0})
}
