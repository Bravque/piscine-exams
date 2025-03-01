package main

import "fmt"

func Gcd(a, b uint) uint {
	// if b == 0 {
	// 	return a
	// }
	// // Recursively call Gcd with b and the remainder of a divided by b
	// return Gcd(b, a%b)
	if a==0 || b==0{
		return 0
	}
	for b!=0{
		a, b = b, a%b
	}
	return a
}

func main() {
	// Test cases
	fmt.Println(Gcd(42, 10))  
	fmt.Println(Gcd(42, 12))  
	fmt.Println(Gcd(14, 77))  
	fmt.Println(Gcd(17, 3))  
}
