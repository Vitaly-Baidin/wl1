package main

import (
	"context"
	"fmt"
	"time"
)

// TODO| Реализовать собственную функцию sleep.

func SleepAfter(duration time.Duration) {
	<-time.After(duration)
}

func SleepTimer(duration time.Duration) {
	<-time.NewTimer(duration).C
}

func SleepTicker(duration time.Duration) {
	<-time.NewTicker(duration).C
}

func SleepContext(duration time.Duration) {
	ctx, _ := context.WithTimeout(context.Background(), duration)
	<-ctx.Done()
}

func main() {
	now := time.Now()
	fmt.Println("Before SleepAfter")
	SleepAfter(2 * time.Second)
	fmt.Println("After SleepAfter", time.Now().Sub(now))

	now = time.Now()
	fmt.Println("Before SleepTimer")
	SleepTimer(2 * time.Second)
	fmt.Println("After SleepTimer", time.Now().Sub(now))

	now = time.Now()
	fmt.Println("Before SleepTicker")
	SleepTicker(2 * time.Second)
	fmt.Println("After SleepTicker", time.Now().Sub(now))

	now = time.Now()
	fmt.Println("Before SleepContext")
	SleepContext(2 * time.Second)
	fmt.Println("After SleepContext", time.Now().Sub(now))
}
