package main

import (
	"fmt"
	"math/big"
)

// TODO| Разработать программу, которая
// TODO| перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	bigA, _ := new(big.Int).SetString("2147483648214748364821474836482147483648", 10)
	bigB, _ := new(big.Int).SetString("4294942949429494294942949429494294942949", 10)

	sum := new(big.Int).Add(bigA, bigB)
	sub := new(big.Int).Sub(bigA, bigB)
	mul := new(big.Int).Mul(bigA, bigB)
	div := new(big.Int).Div(bigA, bigB)

	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)
}
