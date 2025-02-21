package main

import "fmt"

// strcmp compares two strings lexicographically
func strcmp(str1, str2 string) int {
	// Compare each character of both strings
	for i := 0; i < len(str1) && i < len(str2); i++ {
		if str1[i] < str2[i] {
			return -1 // str1 is lexicographically less than str2
		} else if str1[i] > str2[i] {
			return 1 // str1 is lexicographically greater than str2
		}
	}

	// If we've compared all characters and the strings are of different lengths
	if len(str1) < len(str2) {
		return -1 // str1 is shorter, hence less than str2
	} else if len(str1) > len(str2) {
		return 1 // str1 is longer, hence greater than str2
	}

	// If both strings are equal
	return 0
}

func main() {
	// Test cases
	fmt.Println(strcmp("apple", "apple")) // 0 (equal)
	fmt.Println(strcmp("apple", "banana")) // -1 (apple < banana)
	fmt.Println(strcmp("banana", "apple")) // 1 (banana > apple)
	fmt.Println(strcmp("grape", "grapefruit")) // -1 (grape < grapefruit)
}
