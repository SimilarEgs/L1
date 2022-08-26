package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

// we have 2 sets with values
// and one map with key-value int pairs
// task will be completed in 3 cycles

// 1 cycle iterate over first set and write it's value to the declared map, use set value as the key
// and values would be just «1»

// 2 cycle iterate over second set and writes it's values as the key to our previus declared map,
// each key will be increment by 1, so if value from the second set will occur in the map
// this key that occurs in both sets will be increment by 1, and will be equal 2 or more.

//3 cycle prints all keys from the map that equal 2 or more

func main() {

	firstSet := []int{24, 25, 6, 20, 11, 2, 11, 16, 8, 3}
	secondSet := []int{15, 17, 8, 7, 11, 28, 6, 24, 2, 27, 22, 30, 24, 13, 29}

	values := make(map[int]int)

	for _, number := range firstSet {
		values[number]++
	}

	for _, number := range secondSet {
		values[number]++
	}

	for key, val := range values {
		if val >= 2 {
			fmt.Println(key)
		}
	}
}
