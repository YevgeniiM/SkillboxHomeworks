package main

import "fmt"

/*
	Задание 3. Именованные возвращаемые значения
	Что нужно сделать
	Напишите программу, которая на вход получает число, затем с помощью двух функции преобразует его.
	Первая умножает, а вторая прибавляет число, используя именованные возвращаемые значения
*/
func f1(n float64) (a float64) {
	a = n * 10
	return
}
func f2(n float64) (a float64) {
	a = n + 10
	return
}
func main() {
	fmt.Println("Введите число:")
	var number float64
	_, _ = fmt.Scan(&number)
	number = f1(number)
	fmt.Printf("Умножаем на 10 и получаем %v.\n", number)
	number = f2(number)
	fmt.Printf("Прибавляем 10 и получаем %v.\n", number)
}
