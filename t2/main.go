package main

import (
	"fmt"
	"sync"
)

// TODO| Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// TODO| и выведет их квадраты в stdout.

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup // используем WaitGroup для того, чтобы главный поток не завершился раньше остальных

	for _, number := range numbers {
		wg.Add(1)
		go func(number int) { // конкурентно считаем квадраты массива
			fmt.Println(number * number)
			wg.Done()
		}(number)
	}

	wg.Wait()
}
