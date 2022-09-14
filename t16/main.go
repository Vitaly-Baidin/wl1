package main

import "fmt"

// TODO| Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quickSort(arr []int) []int { // сортируем
	return sort(arr, 0, len(arr)-1) // выбор для high длина массива. чтобы использовать только один маркер
}

func sort(arr []int, low, high int) []int {
	if low < high {
		// создаем опорную точку от которой сортируем
		var p int
		arr, p = partition(arr, low, high)
		// сортируем при помощи рекурсии каждый раз сдвигая маркеры
		arr = sort(arr, low, p-1)
		arr = sort(arr, p+1, high)
	}
	return arr
}
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ { // идем по куску слайса
		if arr[j] < pivot { // если меньше опорной точки, меняем местами low и найденный элемент
			arr[i], arr[j] = arr[j], arr[i]
			i++ // двигаем маркер дальше
		}
	}
	arr[i], arr[high] = arr[high], arr[i] // меняем опорную точку
	return arr, i
}

func main() {
	arr := [...]int{1, 2, 3, 12, 151, -4, 0, 86, 0, 11}
	fmt.Println(quickSort(arr[:]))
}
