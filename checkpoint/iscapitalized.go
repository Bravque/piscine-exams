package main

import (
	"fmt"
)

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		if (i == 0 || s[i-1] == ' ') && (s[i] >= 'a' && s[i] <= 'z') {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?")) // false
	fmt.Println(IsCapitalized("Hello How Are You"))  // true
	fmt.Println(IsCapitalized("Whats 4this 100K?"))  // true
	fmt.Println(IsCapitalized("Whatsthis4"))         // true
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))     // true
	fmt.Println(IsCapitalized(""))                   // false
}
