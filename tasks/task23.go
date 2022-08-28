package main

import (
	"fmt"
	"log"
)

// Удалить i-ый элемент из слайса.

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("[Info] slice contains - %v\n", s)
	fmt.Print("[Info] choice index numer to delete: ")

	var i int
	_, err := fmt.Scanln(&i)
	if err != nil {
		log.Fatalf("[Error] %v", err)
	}
	removeSliceElement(s, i)
	fmt.Printf("[Info] slice contains - %v\n", s)

}

func removeSliceElement(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}
