package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		Создать файл только для чтения и проверить, что в него нельзя записать данные
	*/
	file, err := os.Create("Message2.txt")
	if err := os.Chmod("Message2.txt", 0444); err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}
