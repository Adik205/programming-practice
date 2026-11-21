package main

import "fmt"

func main() {
	var total int = 56
	var count int = 13

	var intDiv int = total / count
	var fracDiv float64 = float64(total / count)

	fmt.Println("Целочисленное деление: ", intDiv)
	fmt.Println("Дробное деление: ", fracDiv)

}
