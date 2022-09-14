package main

import (
	"fmt"
	"strings"
)

// TODO| Разработать программу, которая проверяет,
// TODO| что все символы в строке уникальные (true — если уникальные, false etc).
// TODO| Функция проверки должна быть регистронезависимой.
// TODO| Например:
// TODO| abcd — true
// TODO| abCdefAaf — false
// TODO| aabcd — false

func IsUniqueChar(str string) bool {
	m := make(map[rune]struct{})

	for _, ch := range strings.ToLower(str) {
		_, ok := m[ch]
		if ok {
			return false
		}
		m[ch] = struct{}{}
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Printf("%s - %t\n", str1, IsUniqueChar(str1))
	fmt.Printf("%s - %t\n", str2, IsUniqueChar(str2))
	fmt.Printf("%s - %t\n", str3, IsUniqueChar(str3))
}
