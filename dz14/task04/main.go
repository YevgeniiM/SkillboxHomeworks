package main

import "fmt"

/*
Задание 4. Область видимости переменных
Что нужно сделать
Напишите программу, в которой будет 3 функции, попарно использующие разные глобальные переменные.
Функции, должны прибавлять к поданном на вход числу глобальную переменную и возвращать результат.
Затем вызвать по очереди три функции, передавая результат из одной в другую.
*/
var global = 3.14

func f1(n float64) (a float64) {
	a = n + global
	return
}
func f2(n float64) (a float64) {
	a = n + global
	return
}
func f3(n float64) (a float64) {
	a = n + global
	return
}
func main() {
	fmt.Println("Введите число:")
	var number float64
	_, _ = fmt.Scan(&number)
	number = f3(f2(f1(number)))
	fmt.Printf("Трижды прибавив глобальное число 3,14 к исходному получим: %.3f\n", number)
}
