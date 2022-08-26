package main

import "fmt"

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.

// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {

	tmps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tmpGroup := make(map[int][]float64)

	for _, tmp := range tmps {

		keyGap := int(tmp/10) * 10
		tmpGroup[keyGap] = append(tmpGroup[keyGap], tmp)
	}

	for k, v := range tmpGroup {
		fmt.Printf("[Info] temp group %d: %.0f\n", k, v)
	}
}
