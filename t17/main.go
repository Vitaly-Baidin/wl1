package main

import "fmt"

// TODO| Реализовать бинарный поиск встроенными методами языка.

func binarySearch(target int, arr []int) bool {
	low := 0
	high := len(arr) - 1

	if target == arr[low] || target == arr[high] {
		return true
	}

	for low <= high {
		median := (low + high) / 2

		if arr[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) {
		return false
	}

	return true
}

func main() {
	arr := [...]int{-4, 0, 1, 2, 3, 11, 12, 86, 151}
	fmt.Println(binarySearch(0, arr[:]))
}
