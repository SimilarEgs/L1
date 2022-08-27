package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

func main() {

	fmt.Println(reverse("главрыба"))

}

func reverse(input string) string {

	runeString := []rune(input)

	for i, j := 0, len(runeString)-1; i < j; i, j = i+1, j-1 {
		runeString[i], runeString[j] = runeString[j], runeString[i]
	}

	// iterr over string and swap left char with right

	return string(runeString)
}
