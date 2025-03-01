package main

import "fmt"

func StrLen(s string) int {
	return len(s)
}

func main(){
	l := StrLen("Hello world!")
	fmt.Println(l)
}