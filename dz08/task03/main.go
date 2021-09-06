package main

import (
	"fmt"
)

/*
Задание 1:
Пользователь вводит месяц, программа должна вывести, на какое время года
(зима, весна, лето, осень) этот месяц выпадает.
*/

func main() {

	var bills int
	fmt.Println("====================")
	fmt.Println("======================")
	_, _ = fmt.Scanf("%v", &bills)
	for {
		i := 1
		i++
		switch bills {
		case ',':
			fmt.Printf("\n")
		default:

			fmt.Printf("%T", bills)

		}
		if i > 10 {
			break
		}

		/*var lemonadeChange (bills[]int) = false

		switch userDay {
		case "пн":
			fmt.Println("понедельник")
			fallthrough
		case "вт":
			fmt.Println("вторник")
			fallthrough
		case "ср":
			fmt.Println("среда")
			fallthrough
		case "чт":
			fmt.Println("четверг")
			fallthrough
		case "пт":
			fmt.Println("пятница")
		case "сб", "вс":
			fmt.Println("выходные!")
		default:
			fmt.Println("Такого дня нет!")
		}*/
	}
}
