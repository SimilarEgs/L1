package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	stringG := "snow dog sun"

	reverse(stringG)

}

// func shuffle(input string) {

// 	strArr := strings.Split(input, " ")

// 	res := make([]string, len(strArr))
// 	perm := rand.Perm(len(strArr)) // pseudo random, so perm can be [0 1 2]

// 	for i, v := range perm {
// 		res[v] = strArr[i]
// 	}

// 	fmt.Println(res)

// }

func reverse(input string) {

	strArr := strings.Split(input, " ")

	for i := 0; i < len(strArr)/2; i++ {

		strArr[i], strArr[len(strArr)-1-i] = strArr[len(strArr)-1-i], strArr[i]
	}

	fmt.Println(strings.Join(strArr, " "))
}
