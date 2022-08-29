package main

import "fmt"

// Что выведет данная программа и почему?
//
// func someAction(v []int8, b int8) {
//	v[0] = 100
//	v = append(v, b)
//  }
//
//  func main() {
//	var a = []int8{1, 2, 3, 4, 5}
//	someAction(a, 6)
//	fmt.Println(a)
//  }

// Ответ: [1, 2, 3, 4, 5]
// Так происходит, из-за того, что на момент передачи в функцию «someActin» емкость у нашего слайса 5
// а мы хотим добавить в него 6 элемент
// cap и элемент добавиться только внутри самой функции «someAction» так как append вернет новый базовый слайс
//
// Решение: вовзращать новый слайс и присваивать его слайсу в мейне

func someAction(v []int8, b int8) []int8 {
	v[0] = 100
	fmt.Println(cap(v)) // cap 5
	v = append(v, b)
	fmt.Println(cap(v)) // cap 16
	fmt.Println(v)
	return v
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	a = someAction(a, 6)
	fmt.Println(len(a)) // cap 6
	fmt.Println(a)      // [1, 2, 3, 4, 5, 6]
}
