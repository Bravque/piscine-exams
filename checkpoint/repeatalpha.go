package main

import "fmt"

func RepeatAlpha(s string) string {
	var result string

	for _, char := range s {
		var index int
		if char >= 'a' && char <= 'z' {
			index = int(char - 'a') + 1
		} else if char >= 'A' && char <= 'Z' {
			index = int(char - 'A') + 1
		}

		for i := 0; i < index; i++ {
			result += string(char)
		}
	}

	return result 
}

func main() {
	
	fmt.Println(RepeatAlpha("abc"))        
	fmt.Println(RepeatAlpha("choumi."))    
	fmt.Println(RepeatAlpha(""))           
	fmt.Println(RepeatAlpha("abacadaba 01!")) 
}
