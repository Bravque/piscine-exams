package main

import "fmt"

func PrintIf(str string) string {
	if len(str) == 0 {
		return "G" + "\n"
	}
	if len(str) >= 3 {
		return "G" + "\n"
	} else {
		return "Invalid Input" + "\n"
	}
}

func main() {
	fmt.Println(PrintIf("vghdchcjb"))
	fmt.Println(PrintIf("vgh"))
	fmt.Println(PrintIf(""))
	fmt.Println(PrintIf("14"))
}
