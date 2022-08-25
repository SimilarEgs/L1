package main

import (
	"fmt"
	"sync"
	"time"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {

	numbers := []int{2, 4, 6, 8, 10}

	var sumWithMutex int
	var sumWithChannels int

	/////////////////mutex//////////////////////

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			mu.Lock()
			sumWithMutex += num * num
			mu.Unlock()
			wg.Done()
		}(num)
	}
	wg.Wait()

	fmt.Println(sumWithMutex)
	time.Sleep(1 * time.Second)

	/////////////////channels//////////////////////

	ch := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			ch <- num * num
			wg.Done()
		}(num)
	}
	wg.Wait()
	close(ch)

	for number := range ch {
		sumWithChannels += number
	}

	fmt.Println(sumWithChannels)

}
