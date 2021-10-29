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
func main() {
	var str, substr string

	flag.StringVar(&str, "str", "", "set str")
	flag.StringVar(&substr, "substr", "", "set substr")

	flag.Parse()

	fmt.Println(str, substr)
	var result = false
	strRune := []rune(str)
	substrRune := []rune(substr)
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
	fmt.Println(result)
}
