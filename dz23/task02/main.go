package main

import (
	"bufio"
	"fmt"
	"os"
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

func parseTest(sentences [sizeString]string, chars [sizeString]rune) (m [sizeString][sizeString]string) {
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

func enteringValueArrayFromConsole() (output1 [sizeString]string, output2 [sizeString]rune) {
	var value string

	for i := 0; i < sizeString; i++ {
		fmt.Println("Введите строку массива sentences:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		output1[i] = scanner.Text()
	}
	for i := 0; i < sizeRune; i++ {
		fmt.Println("Введите литеру массива chars:")
		_, _ = fmt.Scan(&value)
		//fmt.Println(a)
		for _, output2[i] = range value {
		}
	}
	fmt.Println(output1)
	fmt.Printf("%c\n", output2)
	return output1, output2
}

func main() {
	sentences, chars := enteringValueArrayFromConsole()
	//enteringValueArrayFromConsole()

	m := parseTest(sentences, chars)
	fmt.Println(m)
}
