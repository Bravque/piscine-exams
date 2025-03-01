package main

import "fmt"

func CheckNumber(args string) bool {
	for _, char := range args {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
