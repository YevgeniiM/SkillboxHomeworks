package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Задание 2. Поиск символов в нескольких строках
Что нужно сделать
Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune,
а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово
в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:
func parseTest(sentences []string, chars []rune)
*/
const (
	sizeString = 2
	sizeRune   = 2
)

func parseTest(sentences [sizeString]string, chars [sizeRune]rune) (m [sizeString][sizeRune]string) {
	var a, b string
	for i := 0; i < len(sentences); i++ {
		for j := 0; j < len(chars); j++ {
			a = strings.ToLower(sentences[i])
			b = strings.ToLower(string(chars[j]))
			index := strings.LastIndex(a, b)
			m[i][j] = string(chars[j]) + " position " + strconv.Itoa(index)
		}
	}
	return
}

func main() {
	sentences := [sizeString]string{"qwerty Qwerty", "Sentences string q"}
	chars := [sizeRune]rune{'s', 'Q'}

	m := parseTest(sentences, chars)
	fmt.Println(m)
}
