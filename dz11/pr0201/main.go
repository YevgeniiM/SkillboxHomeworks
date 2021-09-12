package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("вводит строкой число в десятичной системе счисления")
	var number string
	_, _ = fmt.Scan(&number)
	fmt.Println("в какую систему это число перевести:")
	var systemType int
	_, _ = fmt.Scan(&systemType)
	i, err := strconv.ParseInt(number, systemType, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
