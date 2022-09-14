package main

import "fmt"

// TODO| Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func changeBit(target int64, position byte, value byte) int64 {
	if value == 1 {
		return target | (1 << position)
	}

	return target ^ (1 << position)
}

func main() {
	fmt.Println(changeBit(8, 3, 0)) // 1000 > меняем на 3 позиции 1 на 0 > получаем 0000 или 0
	fmt.Println(changeBit(7, 2, 0)) // 111 > меняем на 2 позиции 1 на 0 > получаем 11 или 3
}
