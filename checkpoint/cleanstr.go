package main

import "fmt"

func cleanstr(s string) {
	word := ""
	result := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' {
			word += string(s[i])
		} else if word != "" {
			result += word + " "
			word = ""
		}
	}

	result += word
	fmt.Println(result)
}

func main() {
	cleanstr("you see it's easy to display the same thing")  // "you see it's easy to display the same thing"
	cleanstr(" only    it's  harder   ")                     // "only it's harder"
	cleanstr(" how funny")                                   // "how funny"
	cleanstr("")                                             // ""
}
