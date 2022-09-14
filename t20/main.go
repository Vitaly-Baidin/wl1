package main

import (
	"fmt"
	"strings"
)

// TODO| Разработать программу, которая переворачивает слова в строке.
// TODO| Пример: «snow dog sun — sun dog snow».

func reverseWords(words string) string {
	fields := strings.Fields(words)
	var result []string
	for i := len(fields) - 1; i >= 0; i-- {
		result = append(result, fields[i])
	}

	return strings.Join(result, " ")
}

func main() {
	str1 := "snow dog sun"
	str2 := "asd  asddw    asww"

	str1Reverse := reverseWords(str1)
	str2Reverse := reverseWords(str2)
	fmt.Printf("%s - %s\n", str1, str1Reverse)
	fmt.Printf("%s - %s\n", str2, str2Reverse)
}
