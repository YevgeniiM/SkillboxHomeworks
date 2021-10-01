package main

import "fmt"

func sortBubble(arrayFoSort [6]int) [6]int {
	var k int
	for j := 1; j <= len(arrayFoSort); j++ {
		for i := 0; i < len(arrayFoSort)-j; i++ {
			if arrayFoSort[i] > arrayFoSort[i+1] {
				arrayFoSort[i], arrayFoSort[i+1] = arrayFoSort[i+1], arrayFoSort[i]
				k++ // метка для определения количества замен за проход
			}
		}
		if k == 0 { // если ни одной замены небыло (к не изменился) - значит масив уже отсортирован по возоастанию
			break
		}
		k = 0

	}
	return arrayFoSort
}

func main() {
	array1 := [6]int{29, 4, 55, 30, 15, 5}
	arraySort := sortBubble(array1)
	fmt.Println(arraySort)

}
