package main

import "fmt"

func main() {
	a := 0
	if a > 0 {
		fmt.Println("положительное")
	} else if a < 0 {
		fmt.Println("отрицательное")
	} else {
		fmt.Println("ноль")
	}
}
