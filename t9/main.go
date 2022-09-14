package main

import (
	"fmt"
	"sync"
)

// TODO| Разработать конвейер чисел.
// TODO| Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// TODO| после чего данные из второго канала должны выводиться в stdout.

func main() {
	// как я понял, в один канал данные которые куда-то уходят, в другой канал, дананные уходят в консоль
	var wg sync.WaitGroup
	numbers := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	firstChan := make(chan int, 10)
	secondChan := make(chan int, 10)

	go func() {
		wg.Add(1)
		for _, number := range numbers {
			firstChan <- number
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for _, number := range numbers {
			secondChan <- number * number
		}
		close(secondChan)
		wg.Done()

	}()

	wg.Wait()

	for i := range secondChan {
		fmt.Println(i)
	}
}
