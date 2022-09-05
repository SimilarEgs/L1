package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.

func main() {
						
	numbers := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}

	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			fmt.Printf("square of the number %d: %d\n", num, num*num)
			wg.Done()
		}(num)
	}

	wg.Wait()

}
