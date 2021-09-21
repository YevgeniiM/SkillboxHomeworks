package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	/*
		переписать задачи из урока 2 и 3 на пакет ioutil
	*/
	var b bytes.Buffer

	b.WriteString("Программа выводит сообщения в формате:\n№ строки, дата, сообщение \n")
	in := bufio.NewReader(os.Stdin)
	for j := 1; ; j++ {
		fmt.Println("Введите сообщение:")
		text, _ := in.ReadString('\n')
		if text == "выход\n" {
			b.WriteString("Вышли из программы \n")
			break
		}
		dt := time.Now()
		b.WriteString(fmt.Sprintf("№ %-3d %s    %s", j, dt.Format("01-02-2006 15:04:05"), text))

	}
	fileName := "Message3.txt"
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	factBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Save file:")
	fmt.Println(string(factBytes))
}
