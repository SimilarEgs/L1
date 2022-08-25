package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

func main() {

	mapa := make(map[int]int)
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			mutex.Lock()
			mapa[num] = num
			mutex.Unlock()

		}(i)

	}
	wg.Wait()
	fmt.Println(len(mapa))

	//$ go run -race task07.go - result was without any race(s)
}
