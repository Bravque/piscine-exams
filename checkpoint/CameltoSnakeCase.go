package main

import "fmt"

func CamelToSnakeCase(s string) string {
	for i, char := range s {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return s
		}
		if (s[i] >= 'A' && s[i] <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z') {
			return s
		}

		if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
			return s
		}
	}

	var result string

	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			if i > 0 {
				result += "_"
			}
			result += string(char)
		} else {
			result += string(char)
		}
	}

	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
	fmt.Println(CamelToSnakeCase("Hello@World"))
}
