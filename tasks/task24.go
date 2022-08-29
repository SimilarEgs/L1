package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x, y float32
}

func NewPoint(x, y float32) *Point {
	return &Point{x, y}
}

func main() {
	a := NewPoint(7.1, 4.2)
	b := NewPoint(4.2, 2.3)

	dis := math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2))

	fmt.Println(dis)
}
