package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
Цель задания
Написать программу аналог cat.
Что нужно сделать
Программа должна получать на вход имена двух файлов, необходимо  конкатенировать их содержимое, используя strings.Join.
Что оценивается
    При получении одного файла на входе программа должна печатать его содержимое на экран.
    При получении двух файлов на входе программа соединяет их и печатает содержимое обоих файлов на экран.
    Если программа запущена командой go run firstFile.txt secondFile.txt resultFile.txt, то она должна написать
два соединённых файла в результирующий.
*/

func readFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	factBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return factBytes

}
func makeResultFile(first, second []byte, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(string(first))
	file.WriteString(string(second))
}

func main() {
	//fmt.Println(len(os.Args), os.Args)

	if len(os.Args) == 2 {
		fmt.Println(string(readFile(os.Args[1])))
	} else if len(os.Args) == 3 {
		f := fmt.Sprintf(string(readFile(os.Args[1])) + string(readFile(os.Args[2])))
		f1 := strings.Split(f, "\n")
		fmt.Println(f1)
		fmt.Println("-------------------------------")
		fmt.Println(strings.Join(f1, "\n"))
	} else if len(os.Args) == 4 {
		fmt.Println(os.Args[1], os.Args[2], os.Args[3])
		makeResultFile(readFile(os.Args[1]), readFile(os.Args[2]), os.Args[3])

	} else {
		fmt.Println("Ошибка. Программа должна быть запущена с аргументами в колличестве от 1 до 3")
	}

}
