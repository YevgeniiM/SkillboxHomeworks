package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

//goland:noinspection ALL
func main() {
	//text := ""
	file, err := os.Create("Message.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString("Программа выводит сообщения в формате:\n№ строки, дата, сообщение \n")

	in := bufio.NewReader(os.Stdin)

	for j := 1; ; j++ {

		fmt.Println("Введите сообщение:")
		text, _ := in.ReadString('\n')
		switch text {
		case "выход\n":
			file.WriteString("Вышли из программы \n")
			return
		default:
			dt := time.Now()
			file.WriteString(fmt.Sprintf("№ %-3d %s    %s", j, dt.Format("01-02-2006 15:04:05"), text))
		}

	}

	//file.WriteString()
	//text = "Попытайтесь угадать задуманное число из диапазона от 1 до 100. \n"
	//fmt.Print(text)
	//file.WriteString(text)
	//text = "Программа будет подсказывать введенное число меньше или больше задуманого. \n"
	//fmt.Print(text)
	//file.WriteString(text)
	//
	//for answerNumber != conceiveNumber {
	//	text = "Введите число: \n"
	//	fmt.Print(text)
	//	file.WriteString(text)
	//	fmt.Scan(&answerNumber)
	//	text = "%d \n"
	//	file.WriteString(fmt.Sprintf(text, answerNumber))
	//	if answerNumber < 1 || answerNumber > 100 {
	//		text = "Число должно быть из диапазона от 1 до 100! \n"
	//		fmt.Print(text)
	//		file.WriteString(text)
	//		continue
	//	} else if answerNumber < conceiveNumber {
	//		text = "Задуманное число больше. Попробуйте еше раз! \n"
	//		fmt.Print(text)
	//		file.WriteString(text)
	//	} else if answerNumber > conceiveNumber {
	//		text = "Задуманное число меньше. Попробуйте еше раз! \n"
	//		fmt.Print(text)
	//		file.WriteString(text)
	//	}
	//}
	//text = "Вы угадали! Это число: %d \n"
	//fmt.Printf(text, conceiveNumber)
	//file.WriteString(fmt.Sprintf(text, conceiveNumber))
}
