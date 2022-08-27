package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	count int
	mu    sync.Mutex
	wg    sync.WaitGroup
}

func main() {

	res := Counter{}

	for i := 0; i < 10000; i++ {
		res.wg.Add(1)
		go func(amount int) {
			defer res.wg.Done()
			res.mu.Lock()
			res.count++
			res.mu.Unlock()
		}(i)
	}
	res.wg.Wait()
	fmt.Println(res.count)
}
