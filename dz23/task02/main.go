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
//const (
//	sizeString = 2
//	sizeRune   = 2
//)

func parseTest(sentences []string, chars []rune) [][]string {
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
	return m
}

func enteringValueArrayFromConsole() ([]string, []rune) {
	var value string
	var size int

	fmt.Println("Введите размер массива строк sentences:")
	_, _ = fmt.Scan(&size)
	output1 := make([]string, size)

	for i := 0; i < size; i++ {
		fmt.Println("Введите строку массива sentences:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		output1[i] = scanner.Text()
	}
	fmt.Println("Введите размер массива литер chars:")
	_, _ = fmt.Scan(&size)
	output2 := make([]rune, size)

	for i := 0; i < size; i++ {
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
