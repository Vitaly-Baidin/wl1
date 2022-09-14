package main

import "fmt"

// TODO| Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// TODO| Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// TODO| Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := map[int][]float32{}

	for _, temp := range temperatures {
		step := int(temp) / 10
		value, ok := result[step*10]
		if !ok {
			result[step*10] = []float32{temp}
			continue
		}
		result[step*10] = append(value, temp)
	}
	fmt.Println(result)
}
