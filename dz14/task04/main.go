package main

import "fmt"

/*
Задание 4. Область видимости переменных
Что нужно сделать
Напишите программу, в которой будет 3 функции, попарно использующие разные глобальные переменные.
Функции, должны прибавлять к поданном на вход числу глобальную переменную и возвращать результат.
Затем вызвать по очереди три функции, передавая результат из одной в другую.
*/
var global1, global2, global3 = 3.14, 5.12, 4.65

func f1(n float64) (a float64) {
	a = n + global1 + global2
	fmt.Printf("%.3f\n", a)
	return
}
func f2(n float64) (a float64) {
	a = n + global1 + global3
	fmt.Printf("%.3f\n", a)
	return
}
func f3(n float64) (a float64) {
	a = n + global2 + global3
	fmt.Printf("%.3f\n", a)
	return
}
func main() {
	fmt.Println("Введите число:")
	var number float64
	_, _ = fmt.Scan(&number)
	fmt.Println("================")
	f1(number) // запусуаем функции
	f2(number)
	f3(number)
	number = f3(f2(f1(number))) // вызываем по очереди три функции, передавая результат из одной в другую.
	fmt.Printf("Запустив функции по очереди получим: %.3f\n", number)
}
