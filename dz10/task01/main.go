package main

import (
	"fmt"
	"math"
)

func main() {
	var x, d int
	fmt.Println("Введите для какого значения х посчитать ех:")
	_, _ = fmt.Scan(&x)
	fmt.Println("Ведите с точностью до какого знака выводить результат:")
	_, _ = fmt.Scan(&d)
	epsilon := 1.0 / math.Pow10(d)
	result := 0.0
	prevResult := 0.0
	fact := 1
	n := 0
	for {
		if n > 0 {
			fact *= n
		}
		result += math.Pow(float64(x), float64(n)) / float64(fact)

		if result-prevResult < epsilon {
			fmt.Printf("е в степени %v, с точностью до %v знака, равно %v  \n", x, d, result)
			break
		}
		prevResult = result
		n++
	}
}
