package main

import "fmt"

// counts the alphabetical characters in a string

func CountAlpha(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
}


// Write a function CountAlpha() that takes a string as an argument and returns the number of alphabetic characters in the string.