package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// TODO| Дана последовательность чисел: 2,4,6,8,10.
// TODO| Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

func main() {
	numbers := []int32{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	//var mu sync.Mutex
	var allSum int32

	for _, number := range numbers {
		wg.Add(1)
		go func(number int32) {
			atomic.AddInt32(&allSum, number*number) // атомарно добавляем к числу данные
			//mu.Lock()
			//allSum++
			//mu.Unlock()
			wg.Done()
		}(number)
	}
	wg.Wait()
	fmt.Printf("all sum = %d", allSum)
}
