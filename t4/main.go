package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// TODO| Реализовать постоянную запись данных в канал (главный поток).
// TODO| Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// TODO| Необходима возможность выбора количества воркеров при старте.

// TODO| Программа должна завершаться по нажатию Ctrl+C.
// TODO| Выбрать и обосновать способ завершения работы всех воркеров.

var usageStr = "Usage: reader -cw (int)"
var countWorker int

func main() {
	var wg sync.WaitGroup
	dataChannel := make(chan any, countWorker) // канал для постоянной записи

	wg.Add(countWorker)

	for i := 1; i <= countWorker; i++ {
		go func(id int) {
			for n := range dataChannel {
				fmt.Printf("[worker_id]:%d [msg]:%s\n", id, n) // выводим сообщение с канала
				time.Sleep(500 * time.Millisecond)             // иммитация работы и чтобы консоль не засорялась
			}
			wg.Done()
		}(i)
	}

	signalChan := make(chan os.Signal, 1) // ждем сигнала завершения
	signal.Notify(signalChan, os.Interrupt)
	for {
		select {
		case <-signalChan:
			close(dataChannel) // закрываем канал
			wg.Wait()          // ждем окончания всех горутин
			fmt.Println("All workers close")
			return
		default:
			dataChannel <- "sample data" // отправляем в канал сообщение
		}
	}
}

func init() {
	flag.IntVar(&countWorker, "cw", 1, "count workers") // флаг для запуска -cw [int]
	flag.Usage = func() {
		fmt.Println(usageStr)
	}
	flag.Parse()
}
