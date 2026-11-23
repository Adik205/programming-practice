package main

import "fmt"

func main() {
	var name string = "Хлеб"
	var amount int = 17
	var price float64 = 209.93
	var hasSale bool = true

	fmt.Println("            Информация о товаре:")
	fmt.Println("Название:", name)
	fmt.Printf("Количество: %d шт.\n", amount)
	fmt.Printf("Цена: %.2f тг.\n", price)
	fmt.Println("Имеется скидка: ", hasSale)
}
