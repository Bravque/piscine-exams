package main

import "fmt"

func main() {
	result := ""

	for i := 9; i >= 1; i-- {
		result += fmt.Sprint(i)
	}

	fmt.Println(result)
}
