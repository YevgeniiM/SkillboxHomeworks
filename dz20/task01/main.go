package main

import "fmt"

/*
Задание 1. Подсчёт определителя
Что нужно сделать
Напишите функцию, вычисляющую определитель матрицы размером 3 × 3.
det(A)= a_{11}a_{22}a_{33}+a_{12}a_{23}a_{31}+a_{13}a_{21}a_{32}-a_{13}a_{22}a_{31}-a_{11}a_{23}a_{32}-a_{12}a_{21}a_{33}
*/
const (
	rows = 3
	cols = 3
)

func det(m [rows][cols]int) (detM int) {
	var x, y int
	//
	for j := 0; j < cols; j++ {
		plusDet := 1
		minusDet := 1
		for i := 0; i < rows; i++ {
			y = (i + j) % cols  // по правилу Саррюса присходит эфект переполнения номеров столбцов
			plusDet *= m[i][y]  // перемножаем значения "+ диагонали"
			x = rows - 1 - i    // значения "- диагонали" беме из левого нижнего угла к правому верхнему
			minusDet *= m[x][y] // перемножаем значения "- диагонали"
		}
		detM += plusDet - minusDet
	}
	return
}
func main() {
	matrix := [rows][cols]int{
		{7, 10, 6},
		{2, 15, 8},
		{3, 15, 10},
	}
	det := det(matrix)
	fmt.Println(det)
}
