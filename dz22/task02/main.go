package main

import "fmt"

/*
Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
Что нужно сделать
Заполните упорядоченный массив из 12 элементов и введите число. Необходимо реализовать поиск первого вхождения заданного
числа в массив. Сложность алгоритма должна быть минимальная.
*/
func searchNumber(m []int, a int) (index int) {
	index = -1
	min := 0
	max := len(m) - 1
	i := 0
	for min <= max {
		i++
		fmt.Println(i)
		middle := (max + min) / 2
		switch {
		case m[middle] == a && middle == 0 || m[middle] == a && m[middle-1] < a:
			index = middle
			return
		case m[middle] < a:
			min = middle + 1
		case m[middle] >= a:
			max = middle - 1
		}

	}

	return index
}
func main() {
	m := []int{1, 2, 2, 2, 3, 4, 5, 5, 6, 6, 7, 8, 9, 10, 10, 10, 12}
	fmt.Println(m)
	var number int
	fmt.Println("Введите число")
	_, _ = fmt.Scanln(&number)
	fmt.Println("Первое вхождения заданного числа в массив имеет индекс:", searchNumber(m, number))
}
