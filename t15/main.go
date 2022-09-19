package main

import "fmt"

// TODO| К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// TODO| Приведите корректный пример реализации.

/*
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
*/

func createHugeString(a int) string {
	return "abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc"
}

func someFunc() string {
	// объявление глобальной переменной приведет к негативному последствию, если
	// мы попытаемся до вызова использовать эту переменную она будет равна нулевому значению ""
	// лучше возвращать строку по выполнению некой функции

	v := createHugeString(1 << 10)

	// здесь уже 2 ошибки
	// первая - длина может быть меньше 100 и тогда будет паника
	// вторая - строка это слайс байт, нужны руны
	runes := []rune(v)
	if len(runes) > 100 {
		return string(runes[:100])
	} else {
		return string(runes)
	}
}

func main() {
	fmt.Println(someFunc())
}
