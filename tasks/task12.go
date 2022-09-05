package main

import (
	"fmt"
)

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// Создать для нее собственное множество.

// Множество - почти то же самое, что и хеш-таблица, но без сохранения значения.
// Если нам нужно только быстрая проверка вхождения, то можно снова использовать встроенный map

type void struct{}

var member void

// empty struct doesn’t use any memory

func main() {

	data := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]void)

	for _, v := range data {
		set[v] = member
	}

	for k := range set {
		fmt.Println(k)
	}
}
