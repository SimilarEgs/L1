package main

import "fmt"

// Что такое интерфейсы, как они применяются в Go?

// Интерфейсы - эта конструкция которую используют для описания методов/поведения у объектов.
// Что в свою очередь делает код более гибким и маштабируемым.
// Когда структура(объект) реалезует все методы интерфейса, это означает что она реалезует этот интерфейс.

type Sayer interface {
	Say() string
}

type Dog struct{}
type Cat struct{}
type Pig struct{}

func (d Dog) Say() string {
	return "woof woof"
}
func (p Pig) Say() string {
	return "oink oink"
}
func (c Cat) Say() string {
	return "meow moew"
}

func main() {

	cat := Cat{}
	dog := Dog{}
	pig := Pig{}

	zoo := []Sayer{cat, dog, pig}

	for _, annimal := range zoo {
		fmt.Println(annimal.Say())
	}
}
