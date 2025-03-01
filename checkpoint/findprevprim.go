package main

import "fmt"

func FindPrevPrime(n int) int {
	for ; n > 1; n-- {
		isPrime := true
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return n
		}
	}
	return 0
}

func main() {
	fmt.Println(FindPrevPrime(5)) // 5
	fmt.Println(FindPrevPrime(4)) // 3
}
