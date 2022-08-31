package main

import (
	"fmt"
)

func main() {

	fmt.Println("Зеркальные билеты")
	var digit1, digit2, digit3, digit4, digit5, digit6 int

	min := 100000
	max := 999999
	mirorTickets := 0

	for min <= max {
		digit1 = min % 10
		digit2 = (min / 10) % 10
		digit3 = (min / 100) % 10
		digit4 = (min / 1000) % 10
		digit5 = (min / 10000) % 10
		digit6 = (min / 100000) % 10
		if digit1 == digit6 && digit2 == digit5 && digit3 == digit4 {
			mirorTickets++
		}
		min++
	}
	fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров \nот 100000 до 999999:")
	fmt.Print(mirorTickets)
}
