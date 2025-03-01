package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		return "\n"
	}

	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	return result + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))  // "369\n"
	fmt.Print(ThirdTimeIsACharm(""))           // "\n"
	fmt.Print(ThirdTimeIsACharm("a b c d e f")) // "b e\n"
	fmt.Print(ThirdTimeIsACharm("12"))         // "\n"
}
