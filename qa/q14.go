package main

import "fmt"

// Что выведет данная программа и почему?
//
// func main() {
//	slice := []string{"a", "a"}
//
//	func(slice []string) {
//		slice = append(slice, "a")
//		slice[0] = "b"
//		slice[1] = "b"
//		fmt.Print(slice)
//	}(slice)
//	fmt.Print(slice)
// }

// Ответ: [a a]
// ответ аналогичный q13

func main() {
	slice := []string{"a", "a"}
	fmt.Println(cap(slice)) // cap 2

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(cap(slice)) // cap 4
		fmt.Print(slice)
	}(slice)
	someAction(slice)
	fmt.Print(slice) // [a a]
}

func someAction(slice []string) {
	slice = append(slice, "a")
	slice[0] = "b"
	slice[1] = "b"

}
