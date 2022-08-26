package main

import (
	"fmt"
	"log"
	"strconv"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {

	var num int64
	fmt.Print("[Info] enter desired number: ")
	fmt.Scanln(&num)
	fmt.Printf("[Info] entered number in binary - %b\n", num)

	var iBit int64
	fmt.Print("\n[Info] enter number of the shift bit: ")
	fmt.Scanln(&iBit)

	if iBit < 0 || iBit > 64 {
		log.Fatal("[Error] incorrect bit")
	}

	bitVal := strconv.FormatInt(num^(1<<iBit), 2)
	// XOR(a,b) = 1; only if a != b
	//     else  = 0

	fmt.Printf("[Info] resualt of the program is - %s\n", bitVal)
}
