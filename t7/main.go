package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TODO| Реализовать конкурентную запись данных в map.

type Map struct {
	mu   sync.RWMutex // можно мьютексом, он будет быстрее, но в будущем если чтение с мапы будет намного чаще, то RWMutex побеждает
	data map[int]any
}

func NewMap() *Map {
	return &Map{data: make(map[int]any)}
}

func (m *Map) Set(key int, value any) {
	m.mu.Lock()
	m.data[key] = value
	m.mu.Unlock()
}

func (m *Map) Get(key int) any {
	m.mu.RLock()
	value, ok := m.data[key]
	if !ok {
		value = "not found"
	}
	m.mu.RUnlock()
	return value
}

func main() {
	var wg sync.WaitGroup
	countValue := 10

	myMap := NewMap()
	wg.Add(countValue)
	for i := 0; i < countValue; i++ {
		go func(i int) {
			rand.Seed(time.Now().Unix())
			time.Sleep(1 * time.Millisecond)
			myMap.Set(i, rand.Intn(500))
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < countValue; i++ {
		fmt.Println(myMap.Get(i))
	}
}
