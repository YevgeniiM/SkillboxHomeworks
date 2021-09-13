package main

import (
	"fmt"
	"math"
)

func main() {
	var percent, bankProfit, debit, term, sumPercentPerMonth, factSumPercentPerMonth, sum float64

	fmt.Println("Введите сумму депозита:")
	_, _ = fmt.Scan(&sum)
	fmt.Println("Введите ежемесячный процент капитализаци:")
	_, _ = fmt.Scan(&percent)
	fmt.Println("Введите количество лет на которое делается вклад:")
	_, _ = fmt.Scan(&term)
	debit = sum * 100
	for i := 1; i <= int(term*12); i++ {
		sumPercentPerMonth = debit * percent / 100              //
		factSumPercentPerMonth = math.Floor(sumPercentPerMonth) //float64(int(debit*percent)/100)
		bankProfit += sumPercentPerMonth - factSumPercentPerMonth
		debit += factSumPercentPerMonth
	}
	debit = math.Round(debit*100) / 100
	fmt.Printf("Для %v рублей, %v%% и %v года программа должна вывести %.2f и \n", sum, percent, term, debit/100)
	fmt.Print(bankProfit/100, " (возможно меньше знаков) \n")

}
