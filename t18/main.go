package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// TODO| Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// TODO| По завершению программа должна выводить итоговое значение счетчика.

type Counter interface {
	Add()
	GetValue() int64
}

type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Add() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) GetValue() int64 {
	return c.value
}

type MutexCounter struct {
	value int64
	mu    sync.RWMutex
}

func (m *MutexCounter) Add() {
	m.mu.Lock()
	m.value++
	m.mu.Unlock()
}

func (m *MutexCounter) GetValue() int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.value
}

func test(counter Counter, countGoroutine int) {
	var wg sync.WaitGroup
	wg.Add(countGoroutine * 2)
	for i := 0; i < countGoroutine; i++ {
		go func() {
			counter.Add()
			wg.Done()
		}()
		go func() {
			wg.Done()
		}()
	}
}

func main() {
	var ac Counter = &AtomicCounter{}
	var mc Counter = &MutexCounter{}

	test(ac, 100)
	test(mc, 100)
	fmt.Println(ac.GetValue())
	fmt.Println(mc.GetValue())
}
