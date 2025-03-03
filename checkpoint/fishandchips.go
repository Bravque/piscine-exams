package main

import "fmt"

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}

	if n%2 == 0 && n%3 == 0 {
		return "fish and chips"
	}

	if n%2 == 0 {
		return "fish"
	}

	if n%3 == 0 {
		return "chips"
	}

	return "error: non divisible"
}

func main() {
	fmt.Println(FishAndChips(4))  // Expected: "fish"
	fmt.Println(FishAndChips(9))  // Expected: "chips"
	fmt.Println(FishAndChips(6))  // Expected: "fish and chips"
	fmt.Println(FishAndChips(5))  // Expected: "error: non divisible"
	fmt.Println(FishAndChips(-1)) // Expected: "error: number is negative"
}
