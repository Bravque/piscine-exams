package main

import "fmt"

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	firsthalf := len(str) / 2
	return str[:firsthalf]
}

func main() {
	fmt.Println(RetainFirstHalf("This is the first halfThis is the first half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello world"))
}
