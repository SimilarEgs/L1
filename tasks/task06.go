package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {

	boolExample()
	closedExample()
	contextExample()

}

/////////////////stopping goroutine via signal channel/////////////////
func boolExample() {

	quit := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-quit:
				fmt.Print("stopping goroutine\n\n")
				return
			default:

				fmt.Println("some important work...")
				time.Sleep(time.Second / 2)

			}
		}
	}()

	time.Sleep(2 * time.Second)
	quit <- true

	wg.Wait()
}

/////////////////stopping goroutine via closed channel listener/////////////////
func closedExample() {

	wg := sync.WaitGroup{}

	ch := make(chan int)

	wg.Add(1)
	go func() {

		defer wg.Done()
		for {
			foo, ok := <-ch
			if !ok {
				fmt.Print("channel is closed, stopping goroutine\n\n")
				return
			}
			fmt.Printf("foo value - %d\n", foo)
			time.Sleep(time.Second / 2)
		}
	}()

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)

	wg.Wait()
}

/////////////////stopping goroutine via context/////////////////
func contextExample() {

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan struct{})

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				fmt.Println("cancel() was execute, stopping goroutine")
				return
			default:
				fmt.Println("imitation of some important work...")
			}

			time.Sleep(time.Second / 2)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-ch
	time.Sleep(100 * time.Microsecond)

}
