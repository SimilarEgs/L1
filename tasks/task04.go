package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.

//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {

	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)

	var nWorkers int
	fmt.Print("number of workers: ")
	fmt.Scanln(&nWorkers)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < nWorkers; i++ {
		go consumer(ch)
	}

	func() {
		for {
			select {
			case <-signalChannel:
				fmt.Println("goodbye :)")
				os.Exit(1)
			default:
				go producer(ch)
			}
		}

	}()

}

func producer(ch chan int) {
	for {

		data := rand.Intn(1000)
		ch <- data

	}
}

func consumer(ch chan int) {

	for data := range ch {
		fmt.Printf("data: %d\n", data)
	}

}
