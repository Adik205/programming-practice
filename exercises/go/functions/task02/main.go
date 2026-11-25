package main

import "fmt"

func sumDigits(n int) int {
	sum := 0
	for n != 0 {
		digit := n % 10
		sum += digit
		n /= 10
	}
	return sum
}

func main() {
	num := 10
	fmt.Printf("Сумма цифр числа \"%d\": %d.", num, sumDigits(num))
}
