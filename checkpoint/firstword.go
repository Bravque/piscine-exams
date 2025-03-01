package main

import "fmt"

func FirstWord(s string) string{
	var result string
	start := false

	for _, word:=range s{
		if word == ' ' && !start{
			continue
		}

		if word == ' '{
			break
		}
		result += string(word)
		start = true
	}
	return result
}

func main(){
	fmt.Println(FirstWord("hello there"))
	fmt.Println(FirstWord(""))
	fmt.Println(FirstWord("hello ....... bye"))
	fmt.Println(FirstWord("      hello ....... bye"))
}