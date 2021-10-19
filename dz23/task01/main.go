package main

import (
	"fmt"
)

/*
Задание 1. Чётные и нечётные
Что нужно сделать
Напишите функцию, которая принимает массив чисел, а возвращает два массива:
один из чётных чисел, второй из нечётных.
*/
func makeEvenOddM(m []int) ([]int, []int) {

	sizeOdd := 0
	sizeEven := len(m)

	for i := 0; i < sizeEven; i++ {

		if m[i]%2 == 0 {
			sizeOdd++
			sizeEven = len(m) - sizeOdd
			m[sizeEven], m[i] = m[i], m[sizeEven]
			i--
		}

	}
	//fmt.Println(sizeOdd)

	if sizeOdd == 0 {
		return nil, m
	} else if sizeOdd == len(m) {
		return m, nil
	} else {
		return m[:sizeEven], m[sizeEven:]
	}

}
func main() {
	m := []int{21, 4, 87, 61, 24, 847, 1, 45, 14}
	fmt.Println(m)
	//n := len(m)

	mEven, mOdd := makeEvenOddM(m)
	fmt.Println(mEven)
	fmt.Println(mOdd)
}
