package piscine

import "github.com/01-edu/z01"

// PrintComb3 prints all unique combinations of three different digits in descending order.
func PrintComb3() {
	// Loop through all possible combinations of 3 digits.
	for i := 9; i >= 2; i-- { // i is the first digit
		for j := i - 1; j >= 1; j-- { // j is the second digit
			for k := j - 1; k >= 0; k-- { // k is the third digit
				// Print the combination of digits in the desired format.
				z01.PrintRune(rune(i + '0'))
				z01.PrintRune(rune(j + '0'))
				z01.PrintRune(rune(k + '0'))

				// Check if it's the last combination to avoid trailing comma and space.
				if !(i == 2 && j == 1 && k == 0) {
					z01.PrintRune(',') // Print a comma after the combination
					z01.PrintRune(' ') // Print a space after the comma
				}
			}
		}
	}
}

func main() {
	piscine.PrintComb3()
	z01.PrintRune('\n')
}
