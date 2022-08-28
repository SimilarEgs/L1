package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {

	a := new(big.Float)
	a.SetString("554564645646546546546")

	b := new(big.Float)
	b.SetString("5545123645646546546546")

	res := new(big.Float)

	res.Quo(a, b)

	fmt.Printf("[Info] resualt of division: %f\n", res.Quo(a, b))
	fmt.Printf("[Info] resualt of addition: %f\n", res.Add(a, b))
	fmt.Printf("[Info] resualt of subtraction: %f\n", res.Sub(a, b))
	fmt.Printf("[Info] resualt of multiplication: %f\n", res.Mul(a, b))
}
