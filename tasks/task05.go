package main

import (
	"context"
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

func main() {

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ch := make(chan int)
	go writer(ch)

	for {
		select {
		case <-context.Done():
			fmt.Println("timeout")
			return
		case v := <-ch:
			fmt.Printf("recived data: %d\n", v)
		}
	}
}

func writer(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(time.Second / 2)
	}

}
