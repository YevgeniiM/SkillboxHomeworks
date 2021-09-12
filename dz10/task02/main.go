package main

import (
	"fmt"
	"math"
)

func main() {
	var percent, bankProfit, debit, term, sumPercentPerMonth, factSumPercentPerMonth float64

	fmt.Println("Введите сумму депозита:")
	_, _ = fmt.Scan(&debit)
	fmt.Println("Введите ежемесячный процент капитализаци:")
	_, _ = fmt.Scan(&percent)
	fmt.Println("Введите количество лет на которое делается вклад:")
	_, _ = fmt.Scan(&term)
	//debit *= 100
	for i := 1; i <= int(term*12); i++ {
		sumPercentPerMonth = debit * percent / 100                        //
		factSumPercentPerMonth = math.Floor(sumPercentPerMonth*100) / 100 //float64(int(debit*percent)/100)
		bankProfit += sumPercentPerMonth - factSumPercentPerMonth
		debit += factSumPercentPerMonth
	}
	fmt.Print("К концу срока вклада  сумма депозита составит: ", math.Ceil(debit*100)/100, "грн.\n")
	fmt.Print("Прибыль банка от округлений составит: ", bankProfit, "грн.\n")

}
