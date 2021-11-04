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

func readFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	factBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return factBytes, nil

}
func makeResultFile(first, second []byte, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(string(first))
	if err != nil {
		return err
	}
	_, err = file.WriteString(string(second))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var f1, f2 []byte
	var err error
	//fmt.Println(len(os.Args), os.Args)

	if len(os.Args) == 2 {
		f1, err = readFile(os.Args[1])
		if err != nil {
			fmt.Printf("Не смогли открыть файл: %v\n", err)
		}
		fmt.Println(string(f1))
	} else if len(os.Args) == 3 {
		f1, err = readFile(os.Args[1])
		if err != nil {
			fmt.Printf("Не смогли открыть файл: %v\n", err)
		}
		f2, err = readFile(os.Args[2])
		if err != nil {
			fmt.Printf("Не смогли открыть файл: %v", err)
		}
		f := fmt.Sprintf(string(f1) + string(f2))
		f3 := strings.Split(f, "\n")
		fmt.Println(f3)
		fmt.Println("-------------------------------")
		fmt.Println(strings.Join(f3, "\n"))
	} else if len(os.Args) == 4 {
		f1, err = readFile(os.Args[1])
		if err != nil {
			fmt.Printf("Не смогли открыть файл: %v\n", err)
		}
		f2, err = readFile(os.Args[2])
		if err != nil {
			fmt.Printf("Не смогли открыть файл: %v\n", err)
		}

		err = makeResultFile(f1, f2, os.Args[3])
		if err != nil {
			fmt.Printf("Не смогли создать результирующий файл: %v\n", err)
		} else {
			fmt.Printf("Информация из файлов %v и %v запмсана в результирующий файл %v \n", os.Args[1], os.Args[2], os.Args[3])
		}

	} else {
		fmt.Println("Ошибка. Программа должна быть запущена с аргументами в колличестве от 1 до 3")
	}

}
