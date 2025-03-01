package main

import "fmt"

func ZipString(s string) string {
	result := ""
	
	// Iterate over the string
	for i := 0; i < len(s); i++ {
		count := 1
		// Count the consecutive duplicates
		for i+1 < len(s) && s[i] == s[i+1] {
			count++
			i++
		}
		// Append the count and the character
		result += string(count+'0') + string(s[i])
	}
	
	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}