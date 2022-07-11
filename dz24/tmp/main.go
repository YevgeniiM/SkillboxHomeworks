package main

import "fmt"

/**
Задание 2. Анонимные функции
Что нужно сделать
Напишите анонимную функцию, которая на вход получает массив типа integer, сортирует его пузырьком и переворачивает
(либо сразу сортирует в обратном порядке, как посчитаете нужным).
**/

func main() {

	m := []int{3, 8, 7, 6, 9, 1, 8, 9}
	fmt.Println(m)

	sortInvertedM := func(m []int) /*[]int */ {
		for i := len(m); i > 0; i-- {
			for j := 1; j < i; j++ {
				if m[j-1] < m[j] {
					m[j-1], m[j] = m[j], m[j-1]
				}
			}
		}
		//return m
	}
	sortInvertedM(m)
	fmt.Println("----------------")

	fmt.Println(m)
	//fmt.Println(sortInvertedM(m))

}
