package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	a := 5
	b := 10

	fmt.Printf("a - %d\nb - %d\n\n", a, b)

	a, b = b, a

	fmt.Print("[Info] after swaping values\n\n")

	fmt.Printf("a - %d\nb - %d", a, b)
}
