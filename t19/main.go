package main

import "fmt"

// TODO| Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// TODO| Символы могут быть unicode.

func reverseStr(str string) string {
	runes := []rune(str)
	var result []rune

	for i := len(runes) - 1; i >= 0; i-- {
		result = append(result, runes[i])
	}

	return string(result)
}

func main() {
	source := "главрыба"
	strReverse := reverseStr(source)
	fmt.Printf("%s - %s", source, strReverse)
}
