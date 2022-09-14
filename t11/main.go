package main

import "fmt"

// TODO| Реализовать пересечение двух неупорядоченных множеств.

func intersect(a, b []int) []int {
	dp1 := map[int]struct{}{}
	dp2 := map[int]struct{}{}
	var result []int

	for _, i := range a {
		_, ok := dp1[i]
		if !ok {
			dp1[i] = struct{}{}
		}
	}

	for _, i := range b {
		_, ok := dp1[i]
		if ok {
			_, ok = dp2[i]
			if !ok {
				dp2[i] = struct{}{}
			}
		}
	}

	for i, _ := range dp2 {
		result = append(result, i)
	}

	return result
}

func main() {
	a := []int{2, 4, 8, 6, 11, 24}
	b := []int{8, 7, 24, 24, 24, 33}

	fmt.Println(intersect(a, b))
}
