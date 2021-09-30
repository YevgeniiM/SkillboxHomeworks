package main

import "fmt"

/*
Задание 2. Функция, реверсирующая массив
Что нужно сделать
Напишите функцию, принимающую на вход массив и возвращающую массив, в котором элементы идут
в обратном порядке по сравнению с исходным.
Напишите программу, демонстрирующую работу этого метода.
*/
func arrayRevers(array1 [10]int) (array2 [10]int) {
	for i, r := range array1 {
		reversPosition := len(array2) - 1 - i // определяем обратную позицию
		array2[reversPosition] = r
		fmt.Printf("Записываем значение елемента масива array1[%v] равное %v в елемента масива array2[%v]\n", i, r, reversPosition)
	}
	return
}

func main() {
	array1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arrayRevers(array1))
}
