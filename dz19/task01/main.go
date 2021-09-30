//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
//	BubbleSort(ar)
//	fmt.Println(ar)
//}
//
//func BubbleSort(ar []int) {
//	for i := 0; i < len(ar); i++ {
//		for j := i; j < len(ar); j++ {
//			if ar[i] > ar[j] {
//				swap(ar, i, j)
//			}
//		}
//	}
//}
//
//func swap(ar []int, i, j int) {
//	tmp := ar[i]
//	ar[i] = ar[j]
//	ar[j] = tmp
//}

package main

import (
	"fmt"
)

/*
Задание 1. Слияние отсортированных массивов
Что нужно сделать
Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре
и пять в один массив длиной девять.
*/
func unitArray(a [4]int, b [5]int) (c [9]int) {

	for i := 0; i < len(a)+len(b); i++ {
		if i < len(a) {
			c[i] = a[i]
		} else {
			c[i] = b[i-len(a)]
		}
	}

	return
}

func sortBubble(array3 []int) {
	for i := 0; i < len(array3); i++ {
		for j := i; j < len(array3); j++ {
			if array3[i] > array3[j] {
				swap(array3, i, j)
			}
		}
	}
}

func swap(array3 []int, i, j int) {
	change := array3[i]
	array3[i] = array3[j]
	array3[j] = change
}

func main() {
	array1 := [4]int{1, 4, 8, 12}
	array2 := [5]int{2, 5, 6, 10, 15}

	c := unitArray(array1, array2)

	array3 := []int{c[0], c[1], c[2], c[3], c[4], c[5], c[6], c[7], c[8]}

	sortBubble(array3)
	fmt.Println(array3)
}
