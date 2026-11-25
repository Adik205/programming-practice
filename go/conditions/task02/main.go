package main

import "fmt"

func main() {
	a := 2025
	b := 2022
	c := 2028

	fmt.Print("Максимальное число = ")
	if a >= b && a >= c {
		fmt.Println(a)
	} else if b >= a && b >= c {
		fmt.Println(b)
	} else {
		fmt.Println(c)
	}

}
