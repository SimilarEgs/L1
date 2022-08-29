package main

import "fmt"

// Что выведет данная программа и почему?
//
// func update(p *int) {
//     b := 2
//     p = &b
// }
//
// func main() {
//    var (
//		a = 1
//	    p = &a
//	  )
//	  fmt.Println(*p)
//	  update(p)
//    fmt.Println(*p)
// }

// Ответ: 1
// Почему: внутри функции update мы работает с локальной копией указателя p, а не с значением
// Решение: *p = b

func update(p *int) {
	b := 2
	p = &b
	fmt.Println(p)
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p) // 1
	update(p)       // адрес p втнутри функции - 0xc00001a0f0
	fmt.Println(*p) // 1
	fmt.Println(p)  // 0xc00001a0b8

	rightUpdate(p)  // адрес p втнутри функции - 0xc00001a0b8
	fmt.Println(*p) // нужный нам результат - 2

}

func rightUpdate(p *int) {
	b := 2
	*p = b
	fmt.Println(p)
}
