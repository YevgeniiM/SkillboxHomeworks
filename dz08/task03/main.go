package main

import (
	"fmt"
	"strings"
)

/*
В киоске с лимонадом, каждый лимонад стоит 5 долларов.
Клиенты стоят в очереди, чтобы купить у вас, и заказывают по одному (в порядке, указанном в счетах).
Каждый покупатель может купить только один лимонад и заплатить купюрой на 5, 10 или 20 долларов.
Вы должны дать каждому покупателю сдачу с его купюры.
*/

func main() {
	var bill5, bill10, bill20 int
	var bills string
	var lemonadeChange bool
	fmt.Println("В киоске с лимонадом, каждый лимонад стоит 5 долларов.")

		fmt.Println("Введите, через запятую, какие купюры у покупателей (в порядке, указанном в счетах):")
		_, _ = fmt.Scan(&bills)
		bill := strings.Split(bills, ",")
		for i:= 0; i < len(bill); i++ {
			switch {
			case bill [i]== "5":
				bill5++
				lemonadeChange = true
			case bill [i]== "10":
				bill10++
				bill5 --
				if bill5 >= 0 {
					lemonadeChange = true
				} else {
					lemonadeChange = false
				}

			case bill [i]== "20":
				bill20++
				if bill5 - 1 >= 0 && bill10  - 1 >= 0 {
					bill5 --
					bill10 --
					lemonadeChange = true
				} else if bill5 - 3 >= 0{
					bill5 -= 3
					lemonadeChange = true
				} else {
					lemonadeChange = false
				}
			default:
				lemonadeChange = false
				fmt.Println(lemonadeChange, "bills only 5,10,20")
				return
			}
		}
	fmt.Println(lemonadeChange)
		summ := bill5*5 + bill10*10 + bill20*20
	fmt.Print("Мы заработали: ",summ, "$\n")
}

