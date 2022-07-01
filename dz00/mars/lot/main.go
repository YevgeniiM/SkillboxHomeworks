package main

import (
	"fmt"
)

/*
   Задание 2. Поиск символов в нескольких строках
   Что нужно сделать
   Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune,
   а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово
   в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:
   func parseTest(sentences []string, chars []rune)
*/
//const (
//	sizeString = 2
//	sizeRune   = 2
//)

/*func newLot(sentences []string, chars []rune) {
	var a, b string
	var lo, hi int
	s := make([]string, len(chars)*len(chars))

	m := make([][]string, len(sentences))

	for i := 0; i < len(sentences); i++ {
		a = strings.ToLower(sentences[i])
		for j := 0; j < len(chars); j++ {
			hi = j + len(chars)*i
			b = strings.ToLower(string(chars[j]))
			index := strings.LastIndex(a, b)
			if index >= 0 {
				s[hi] = "'" + string(chars[j]) + "'" + " position " + strconv.Itoa(index)
			}
		}
		m[i] = s[lo : hi+1]
		lo += len(chars)
	}
	return
}
*/
func oldLot() [][]int {
	//var value int
	var size, maxNumber, timeLot int

	fmt.Println("Введите колличество шаров в розыграше:")
	_, _ = fmt.Scan(&maxNumber)

	fmt.Println("Введите колличество розыграшей:")
	_, _ = fmt.Scan(&timeLot)

	fmt.Println("Введите колличество номеров в розыграше:")
	_, _ = fmt.Scan(&size)

	r := make([]int, size)
	output2 := make([][]int, timeLot*size)

	for i := 0; i < timeLot; i++ {

		fmt.Println("Введите результат розыграша:")
		_, _ = fmt.Scan(&r[0], &r[1], &r[2], &r[3])
		fmt.Println(r)

		for j := 1; j < 4; j++ {
			output2[i][j] = int(r[j])
		}

		/*}
		for i := 0; i < size; i++ {
			fmt.Println("Введите литеру массива chars:")
			_, _ = fmt.Scan(&value)
			//fmt.Println(a)
			for _, output2[i] = range value {
			}*/
	}
	//fmt.Println(output2)
	//fmt.Printf("%c\n", output2)
	return output2
}

func main() {
	//sentences, chars := enteringValueArrayFromConsole()
	//enteringValueArrayFromConsole()

	m := oldLot()
	fmt.Println(m)
}
