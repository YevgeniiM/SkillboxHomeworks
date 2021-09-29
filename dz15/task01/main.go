package main

import "fmt"

/*
Задание 1. Подсчёт чётных и нечётных чисел в массиве
Что нужно сделать
Разработайте программу, позволяющую ввести 10 целых чисел, а затем вывести из них количество чётных и нечётных чисел.
Для ввода и подсчёта используйте разные циклы.
*/
func main() {
	var masiv [10]int
	var oddNumber, evenNumber int

	fmt.Println("Необходимо ввести 10 целых чисел")

	for i := range masiv {
		fmt.Println("ведите число номер", i+1)
		_, _ = fmt.Scan(&masiv[i])
	}

	for _, number := range masiv {
		if number%2 == 0 {
			oddNumber++
		} else {
			evenNumber++
		}
	}
	fmt.Printf("Среди введённых чисел четных %v и нечетных %v", oddNumber, evenNumber)
}
