package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой. Например: abcd — true abCdefAaf — false aabcd — false

func main() {

	var inputString string

	fmt.Print("[Info] enter ur string: ")
	fmt.Scanln(&inputString)

	runString := []rune(inputString)

	contains := make(map[string]int)

	for _, c := range string(runString) {
		contains[strings.ToLower(string(c))]++
	}

	for _, v := range contains {
		if v > 1 {
			fmt.Println("[Info] not unique")
			return
		}
	}

	fmt.Println("[Info] string unique")

}
