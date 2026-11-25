package main

import "fmt"

func main() {
	balance := 10000
	fmt.Printf("Баланс - %d тг.\n", balance)
	balance += 3500
	fmt.Printf("Пополнение на 3 500 тг. Баланс - %d тг.\n", balance)
	balance -= 4200
	fmt.Printf("Покупка на 4 200 тг. Баланс - %d тг.\n", balance)
	balance -= 1800
	fmt.Printf("Покупка на 1 800 тг. Баланс - %d тг.\n", balance)
	balance += 700
	fmt.Printf("Возврат на 700 тг. Баланс - %d тг.\n", balance)
}
