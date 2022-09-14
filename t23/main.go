package main

import "fmt"

// TODO| Удалить i-ый элемент из слайса.

func DeleteElement(source []int, index int) []int {
	var result []int
	for i := 0; i < len(source); i++ {
		if i == index {
			continue
		}
		result = append(result, source[i])
	}
	return result
	//return append(source[:index], source[index+1:]...) плохой вариант, т.к. меняет исходный массив
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println("source array before delete:", arr)

	newArr := DeleteElement(arr, 2)
	fmt.Println("source array after delete:", arr)
	fmt.Println("new array:", newArr)

	newArr[0] = 2222
	fmt.Println("source array after changed new array:", arr)
	fmt.Println("new array:", newArr)
}
