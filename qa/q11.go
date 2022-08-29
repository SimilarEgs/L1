package main

// Что выведет данная программа и почему?
//
// func main() {
// 	 wg := sync.WaitGroup{}
// 	 for i := 0; i < 5; i++ {
// 	    wg.Add(1)
// 	    go func(wg sync.WaitGroup, i int) {
// 	 	  fmt.Println(i)
// 	 	  wg.Done()
// 	    }(wg, i)
// 	 }
// 	 wg.Wait()
// 	 fmt.Println("exit")
//  }

// Ответ: ошибка - deadlock
// Решение: Передавать вейт-группу по указателю
// Почему: горутина работает со значением wg а не с указателем на это значение,
// то есть мы работает с копией wg: «func passes lock by value: sync.WaitGroup contains sync.noCopycopylocks»
