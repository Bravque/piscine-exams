package main

import "fmt"

func LastRune(s string) rune {
	var last rune

	// Iterate over the string from the end towards the beginning
	for i := len(s) - 1; i >= 0; i-- {
		char := rune(s[i])
		if char != ' ' { // Skip spaces
			last = char
			break
		}
	}
	return last
}

func main() {
	fmt.Println(string(LastRune("Hello!")))   // Output: !
	fmt.Println(string(LastRune("Salut!")))   // Output: !
	fmt.Println(string(LastRune("Ola!   ")))  // Output: !
	fmt.Println(string(LastRune("   Test "))) // Output: t
	fmt.Println(string(LastRune("   ")))      // Output: (no output)
}
