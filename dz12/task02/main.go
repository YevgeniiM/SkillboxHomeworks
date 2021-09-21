package main

import (
	"fmt"
	"os"
)

/*
Урок №3 интерфейс io.Reader
Написать программу, которая полностью вычитывала бы файл из предыдущего домашнего задания без использования ioutil
Подсказка: для получения размера файла у файла есть метод Stat(), который возвращает информацию о файле и ошибку.
*/
func main() {
	f, err := os.Open("Log.txt")
	if err != nil {
		fmt.Println("Not open file", err)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		panic(err)
		return
	}
	size := fi.Size()
	buf := make([]byte, size)
	_, err = f.Read(buf)
	if err != nil {
		fmt.Println("Not red file", err)
		return
	}
	fmt.Println(string(buf))
}
