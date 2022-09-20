package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TODO| Реализовать все возможные способы остановки выполнения горутины.

func main() {
	var wg sync.WaitGroup

	ctx := context.Background()

	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	fmt.Println("close goroutine (WithTimeout 5 sec)")
	wg.Add(1)
	go func() {
		<-ctxTimeout.Done()
		wg.Done()
	}()
	wg.Wait()
	cancel()

	ctxDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	fmt.Println("close goroutine (WithDeadline 5 sec)")
	wg.Add(1)
	go func() {
		<-ctxDeadline.Done()
		wg.Done()
	}()
	wg.Wait()
	cancel()

	ctxCancel, cancel := context.WithCancel(ctx)
	fmt.Println("close goroutine (WithCancel)")
	wg.Add(1)
	go func() {
		go func() {
			rand.Seed(time.Now().Unix())
			for {
				intn := rand.Intn(100)
				if intn > 50 {
					cancel()
					return
				}
			}
		}()
		<-ctxCancel.Done()
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("close goroutine (channel)")
	wg.Add(1)
	channel := make(chan int)
	go func() {
		rand.Seed(time.Now().Unix())
		for {
			intn := rand.Intn(100)
			if intn > 50 {
				channel <- 1
				close(channel)
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case <-channel:
				wg.Done()
				fmt.Println("go close")
				return
			default:
				fmt.Println("go no close")
			}
		}
	}()

	wg.Wait()
}
