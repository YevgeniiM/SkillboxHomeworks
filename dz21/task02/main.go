package main

import "fmt"

/*
Задание 2. Анонимные функции
Что нужно сделать
Напишите функцию, которая на вход принимает функцию вида A func (int, int) int, а внутри оборачивает
и вызывает её при выходе (через defer).
Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций могут быть любыми, но главное,
чтобы все три выполняли разное действие.
*/
func outDefer(x, y int, a func(int, int) int) int {
	defer a(x, y)
	x = x * 2
	y = y + 10
	return a(x, y)
}

func main() {
	x := 10
	y := 5
	fmt.Println(outDefer(x, y, func(x, y int) int { return x + y }))
	fmt.Println(outDefer(x, y, func(x, y int) int { return x / y }))
	fmt.Println(outDefer(x, y, func(x, y int) int { return x * y }))

}
