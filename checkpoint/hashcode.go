package main

import "fmt"

func Hashcode(dec string) string {
	var result string
	for _, char := range dec {
		// Calculate hash for each character
		hash := (int(char) + len(dec)) % 127

		// Adjust hash if it's outside the valid range
		if hash < 32 || hash > 127 {
			hash += 33
		}

		// Convert hash to string directly by casting to a rune
		result += string(rune(hash))
	}

	return result
}

func main() {
	fmt.Println(Hashcode("A"))           
	fmt.Println(Hashcode("AB"))         
	fmt.Println(Hashcode("BAC"))         
	fmt.Println(Hashcode("Hello World")) 
	fmt.Println(Hashcode(""))

}
