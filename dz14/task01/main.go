package main

import "fmt"

/*
Задание 1. Функция возвращающая результат
Что нужно сделать
Напишите функцию, которая на вход получает число и возвращает  true, если число четное, и false, если нечетное.
Советы и рекомендации
Программа запрашивает у пользователя или генерирует случайное число, передает в функцию в качестве аргумента и
выводит в консоль результат ее работы.
*/
func f(x int) bool {
	return x%2 == 0
}

func main() {
	fmt.Println("Программа по определению четное число или нет")
	fmt.Println("Программа выведет true, если число четное, и false, если нечетное")
	fmt.Println("Введите число:")
	var a int
	_, _ = fmt.Scan(&a)
	fmt.Println(f(a))
}