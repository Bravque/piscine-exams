package main

import "fmt"

func main(){
	firstcharacter("mnsion")
}

func firstcharacter(str string) {
	for _, r:=range str{
		if r == 'a'{
			fmt.Print(string(r))
			fmt.Println("\n")
		}
	}
	fmt.Println()
}

/*Write a program that takes a string, and displays the first 'a' character it
encounters in it, followed by a newline. If there are no 'a' characters in the
string, the program just writes a newline. If the number of parameters is not
1, the program displays 'a' followed by a newline.*/
