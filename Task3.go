package main

import (
	"fmt"
)

func main() {

	var h int

	fmt.Println("Вывод ёлочки.")

	fmt.Println("Введите высоту:")
	fmt.Scan(&h)

	for i := 0; i < h; i++ {
		for j := 0; j < h-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
