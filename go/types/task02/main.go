package main

import "fmt"

func main() {
	minutes := 135

	fullHours := minutes / 60
	remainMinutes := minutes % 60

	fmt.Printf("%d минут = %d часов %d минут.", minutes, fullHours, remainMinutes)
}
