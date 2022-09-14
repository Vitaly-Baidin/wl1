package main

import "fmt"

// TODO| Поменять местами два числа без создания временной переменной.

func change(a, b int) (int, int) {
	b = b - a
	a = a + b
	b = a - b
	return a, b
}

func main() {
	fmt.Println(change(12, 4))
	fmt.Println(change(5, 11))
	fmt.Println(change(-7, -44))

	a := 13
	b := 224
	fmt.Println(a, b)

	a, b = b, a // один из вариантов, без математики
	fmt.Println(a, b)
}
