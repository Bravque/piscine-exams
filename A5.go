// Write a program that displays the alphabet, with even letters in uppercase, and
// odd letters in lowercase, followed by a newline.

package main

import "fmt"

func main() {
	for i := 0; i < 26; i++ {
		if i%2 == 0 {
			fmt.Print(string('A' + i)) // Uppercase
		} else {
			fmt.Print(string('a' + i)) // Lowercase
		}
	}
	fmt.Println() // Newline
}
