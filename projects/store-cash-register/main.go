package main

import (
	"fmt"
	"os"
	"strconv"
)

func readPositiveInt(prompt string) (int, error) {
	for {
		fmt.Print(prompt)

		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			return 0, err
		}

		value, err := strconv.Atoi(input)
		if err != nil || value <= 0 {
			fmt.Println("Введите целое число больше нуля: ")
			continue
		}

		return value, nil
	}
}

func main() {
	sum := 0
	for {
		var price int
		var amount int
		var nextProduct string

		price, err := readPositiveInt("Цена товара, ₸: ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Не удалось прочитать центу товара: ", err)
			os.Exit(1)
		}

		amount, err = readPositiveInt("Количество, шт.: ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Не удалось прочитать количество товара: ", err)
			os.Exit(1)
		}

		sum += price * amount

		fmt.Print("Хотите добавить следующий товар: ")
		for nextProduct != "да" && nextProduct != "нет" {
			_, err := fmt.Scan(&nextProduct)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Не удалось прочитать ответ: ", err)
				os.Exit(1)
			}
			if nextProduct != "да" && nextProduct != "нет" {
				fmt.Print("Введите \"да\" или \"нет\": ")
			}
		}

		if nextProduct == "нет" {
			break
		}
	}

}
