package main

import (
	"flag"
	"fmt"
)

/**
Задача
Цель задания
Написать программу для нахождения подстроки в кириллической подстроке. Программа должна запускаться с помощью команды:
 go run main.go --str "строка для поиска" --substr "поиска"
Для реализации такой работы с флагами воспользуйтесь пакетом flags, а для поиска подстроки в строке вам понадобятся руны.
**/
func stringComparison(a string, b string) (result bool) {
	strRune := []rune(a)
	substrRune := []rune(b)

	for i, r := range strRune {
		if substrRune[0] == r {
			result = true
			for j, r := range substrRune {
				if strRune[i+j] != r {
					result = false
					break
				}

			}
			if result {
				break
			}
		}
	}
	return
}

func main() {
	var str, substr string

	flag.StringVar(&str, "str", "", "set str")
	flag.StringVar(&substr, "substr", "", "set substr")

	flag.Parse()

	fmt.Println(str, substr)

	if str == "" || substr == "" {
		fmt.Println("Error! нельзя сравнивать пустые строки")
		return
	}
	fmt.Println(stringComparison(str, substr))
}
