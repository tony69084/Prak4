package main

import (
	"fmt"
)

func main() {

	// h - высота, d - ширина
	var h, d int

	fmt.Println("Шахматная доска.")

	fmt.Println("Введите ширину:")
	fmt.Scan(&d)

	fmt.Println("Введите высоту:")
	fmt.Scan(&h)

	for i := 0; i < h; i++ {
		for j := 0; j < d; j++ {
			if (j+i)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}

}
