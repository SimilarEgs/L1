package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// func someFunc() {
//     v := createHugeString(1 << 10)
//     justString = v[:100]
// }
//
// func main() {
//     someFunc()
// }

// HugeString - it's a massive of bytes, so if we cut this string using the example above
// we may face next problem - some characters in this string may consist more than one byte
// and that cause problem because the result of cut will not be correct

// Also, in this section we allocate more memory than we need: «justString = v[:100]»
// allocated unused memory for this chunk of code wll exist until this variable is no longer referenced.

// solution:
// Convert «HugeString» to slice of rune,
// each element of the rune consists of Unicode symbol and is equal to 1 bit

// func main() {

// 	var justString string
// 	v := createHugeString(150)
// 	justString = v[:100]
// 	fmt.Println(justString) //----------------> in total 50 symbols without spaces instead of desired 100

// }

const (
	charNum = 200
)

func main() {

	str := idiomaticWay()
	fmt.Println(str) //----------------> in total 100 symbols
}

func createHugeString(numC int) string {
	return strings.Repeat("ЭЫ)Юы ", numC)
}

func idiomaticWay() string {

	var justString string

	hugeString := createHugeString(charNum)

	runeHugeString := []rune(hugeString)[:100]

	slicedRuneHugeString := make([]rune, len(runeHugeString))

	copy(slicedRuneHugeString, runeHugeString)

	justString = string(slicedRuneHugeString)

	return justString
}
