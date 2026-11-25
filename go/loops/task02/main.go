package main

import "fmt"

func main() {
	n := 32432
	digits := 1

	for remaining := n; remaining >= 10; remaining /= 10 {
		digits++
	}

	fmt.Printf("У числа %d есть %d цифр.", n, digits)
}
