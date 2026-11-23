package main

import "fmt"

func main() {
	firstRate := 4
	secondRate := 4
	thirdRate := 5
	sumOfRates := firstRate + secondRate + thirdRate
	arithMean := float64(sumOfRates) / 3

	fmt.Printf("Среднее арифметическое всех оценок - %.3f.", arithMean)
}
