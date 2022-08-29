package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Какой самый эффективный способ конкатенации строк?

// Ответ: используя буффер пакета «bytes» или «strings.Builder»
// Почему: этот метод не ведет к созданию новых строк
// в go строки это тип read-only, и любая манипуляция с ними создает новую строку, что не есть эффективно
// Так же если известна итоговая длинна строки, можно воспользоваться функцией «Copy()», что тоже будет эффективно

func main() {
	var buffer bytes.Buffer

	var sb strings.Builder

	for i := 0; i < 1000; i++ {
		buffer.WriteString("a")
	}

	for i := 0; i < 1000; i++ {
		sb.WriteString("a")
	}

	fmt.Println(buffer.String())
}
