package main

import "fmt"

func main() {
	totalSeconds := 12849
	hours := totalSeconds / 3600
	minutes := (totalSeconds / 60) % 60
	seconds := totalSeconds % 60

	fmt.Printf("%d секунд = %d часов %d минут %d секунд.", totalSeconds, hours, minutes, seconds)

}
