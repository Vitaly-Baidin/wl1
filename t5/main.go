package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

// TODO| Разработать программу, которая будет последовательно отправлять значения в канал,
// TODO| а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

var usageStr = "Usage: reader -t (int)"
var timeout int

func main() {
	dataChannel := make(chan any)
	ctx := context.Background()
	ctxTimeout, _ := context.WithTimeout(ctx, time.Duration(timeout)*time.Second) // завершаем работу при помощи контекста

	go func() {
		for {
			time.Sleep(1 * time.Second)
			dataChannel <- "sample data"
		}
	}()

	for {
		select {
		case <-ctxTimeout.Done(): // как только время в контексте выйдет, он отправит сигнал завершения
			close(dataChannel)
			fmt.Println("timeout")
			return
		case msg := <-dataChannel:
			fmt.Println("[msg]:", msg) // иначе выводим сообщения
		}
	}
}

func init() {
	flag.IntVar(&timeout, "t", 0, "count workers")
	flag.Usage = func() {
		fmt.Println(usageStr)
	}
	flag.Parse()
}
