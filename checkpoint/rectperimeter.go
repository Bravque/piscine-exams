package main

import "fmt"

// function to calculate perimeter

func RectPerimeter(w, h int) int {
	if h < 0 || w < 0 {
		return -1
	}
	return 2 * (w + h)
}

func main() {
	fmt.Println(RectPerimeter(-10, 2))
}
