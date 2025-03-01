package main

import "fmt"

func PrintMemory(arr [10]byte) {
	// Print hexadecimal representation
	for _, b := range arr {
		fmt.Printf("%02x ", b)
	}
	fmt.Println()

	// Print ASCII representation
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			fmt.Printf("%c", b)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*', 0, 0})
}
