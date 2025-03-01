package main

import "fmt"

func Compare(a, b string) int {
	// Compare character by character
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	// If strings are the same up to the minimum length, compare their lengths
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}

	return 0
}

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))  // 0
	fmt.Println(Compare("Salut!", "lut!"))    // -1
	fmt.Println(Compare("Ola!", "Ol"))        // 1
}
