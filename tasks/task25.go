package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	sleepImpostor(5)
	fmt.Println(123)
}

func sleepImpostor(d time.Duration) {
	<-time.After(time.Second * d)
}
