package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {

	wg := sync.WaitGroup{}

	numbers := []int{2, 4, 6, 8, 10}

	ch1 := make(chan int, len(numbers))
	ch2 := make(chan int, len(numbers))

	wg.Add(3)
	// 1 gorutina - read data from arr
	// 2 gorutina - processes data from the first chan
	// 3 gorutina - writes data from the second chan

	go func() {
		defer wg.Done()

		for _, num := range numbers {
			ch1 <- num
		}
		close(ch1)
	}()

	go func() {
		defer wg.Done()

		for num := range ch1 {
			ch2 <- num * 2
		}
		close(ch2)
	}()

	go func() {
		defer wg.Done()

		for num := range ch2 {
			time.Sleep(time.Second / 3)
			fmt.Fprintln(os.Stdout, num)
		}

	}()

	wg.Wait()

}
