package main

import "fmt"

func rectangleArea(width, height int) int {
	return width * height
}
func main() {
	a := 4
	b := 6
	fmt.Println("S = ", rectangleArea(a, b))
}
