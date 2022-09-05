package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// O(log n) - сложность

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	arr := []int{29, 10, 9, 17, 28, 4, 8, 20, 6, 1}
	quicksort(arr)
	fmt.Println(arr)

}

func quicksort(arr []int) {

	if len(arr) <= 1 {
		return
	}

	left, right := 0, len(arr)-1

	pivot := rand.Int() % len(arr)

	arr[right], arr[pivot] = arr[pivot], arr[right]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quicksort(arr[:left])
	quicksort(arr[left+1:])

}
