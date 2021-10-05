package main

import (
	"fmt"
	"math"
)

/*
Задание 1. Расчёт по формуле
Что нужно сделать
Напишите функцию, производящую следующие вычисления.
S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.
Тип S должен быть во float32.
*/
func s(x int16, y uint8, z float32) float32 {
	result := 2*float32(x) + float32(math.Pow(float64(y), 2)) - 3/z
	return result
}

func main() {
	var x int16 = 10
	var y uint8 = 14
	var z float32 = 30

	resultS := s(x, y, z)
	fmt.Println(resultS)
}
