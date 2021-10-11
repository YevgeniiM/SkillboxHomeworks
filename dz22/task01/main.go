package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*Задание 1. Подсчёт чётных и нечётных чисел в массиве
Что нужно сделать
Заполните массив неупорядоченными числами на основе генератора случайных чисел. Введите число.
Программа должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого.
При отсутствии введённого числа в массиве — вывести 0. Для удобства проверки реализуйте вывод массива на экран.
*/
const n = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	var m [n]int
	for i := 0; i < n; i++ {
		m[i] = rand.Intn(100)
	}
	fmt.Println(m)
	var number int
	numberAfter := -1
	fmt.Println("Введите число:")
	_, _ = fmt.Scanln(&number)

	for i := 0; i < len(m); i++ {
		if number == m[i] {
			numberAfter = len(m) - 1 - i
		}
	}
	if numberAfter == -1 {
		fmt.Printf("Числа %d в массиве нет", number)

	} else {

		fmt.Printf("После числа %d в массиве находится %d чисел", number, numberAfter)
	}
}
