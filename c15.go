package main

import "fmt"

// isPowerOfTwo checks if the given number is a power of 2
func isPowerOfTwo(n int) int {
	// If n is greater than 0 and n & (n - 1) == 0, it's a power of 2
	if n > 0 && (n&(n-1)) == 0 {
		return 1
	}
	return 0
}

func main() {
	// Test cases
	fmt.Println(isPowerOfTwo(16))  // 1 (16 is 2^4)
	fmt.Println(isPowerOfTwo(18))  // 0 (18 is not a power of 2)
	fmt.Println(isPowerOfTwo(32))  // 1 (32 is 2^5)
	fmt.Println(isPowerOfTwo(1))   // 1 (1 is 2^0)
	fmt.Println(isPowerOfTwo(0))   // 0 (0 is not a power of 2)
}
