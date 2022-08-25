package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	name string
}

type Action struct {
	Human
}

func (h Human) sing() string {
	return fmt.Sprintf("I'm only %s after all", h.name)
}

func main() {
	human01 := Human{name: "Human"}
	action01 := Action{Human: human01}

	resH := human01.sing()
	resA := action01.Human.sing() + " don't put your blame on me"

	fmt.Println(resH) // I'm only Human after all
	fmt.Println(resA) // I'm only Human after all don't put your blame on me
}
