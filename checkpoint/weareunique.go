package main

import (
	"fmt"
)

func UniqueCharCount(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	uniqueCount := 0

	// Check unique characters in str1
	for _, c := range str1 {
		if !contains(str2, c) {
			uniqueCount++
		}
	}

	// Check unique characters in str2
	for _, c := range str2 {
		if !contains(str1, c) {
			uniqueCount++
		}
	}

	return uniqueCount
}

// Helper function to check if a character is in a string
func contains(s string, ch rune) bool {
	for _, c := range s {
		if c == ch {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(UniqueCharCount("foo", "boo")) // Output: 2
	fmt.Println(UniqueCharCount("", ""))       // Output: -1
	fmt.Println(UniqueCharCount("abc", "def")) // Output: 6
}
