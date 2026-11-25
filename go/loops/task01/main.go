package main

import "fmt"

func main() {
	sumOfEvenNums := 0
	n := 9

	for i := 2; i <= n; i += 2 {
		sumOfEvenNums += i
	}

	fmt.Println(sumOfEvenNums)
}
