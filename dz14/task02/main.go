package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Задание 2. Функция возвращающая несколько значений
Что нужно сделать
Напишите программу, которая с помощью функции генерирует 3 случайные точки в двумерном пространстве (две координаты),
а затем с помощью другой функции преобразует эти координаты по формулам: x = 2 * x + 10, y = -3 * y - 5.
*/
func randomPoint(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func changePoint(x, y int) (int, int) {
	x = 2*x + 10
	y = -3*y - 5
	return x, y
}

func main() {
	fmt.Println("Введите вехнюю границу диапазона (0-n) генерации случайных координат для точек:")
	var n int
	_, _ = fmt.Scan(&n)
	fmt.Println("Введите сколько случайных точек необходимо сгенерировать:")
	var numberPoint int
	_, _ = fmt.Scan(&numberPoint)
	var x, y int

	for i := 0; i < numberPoint; i++ {
		x = randomPoint(n) // генерируем координаты
		y = randomPoint(n)
		fmt.Printf("Исходные координаты: x =%4d y =%4d\n", x, y)
		x, y = changePoint(x, y)
		fmt.Printf("Новые координаты:    x =%4d y =%4d\n", x, y)
	}
}
