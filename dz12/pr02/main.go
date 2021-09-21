package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//goland:noinspection ALL
func main() {
	text := ""
	file, err := os.Create("Log.txt")
	if err != nil {
		return
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())
	conceiveNumber := 15
	//conceiveNumber := rand.Intn(101)
	text = "задуманное число %d \n"
	file.WriteString(fmt.Sprintf(text, conceiveNumber))

	var answerNumber int
	text = "Программа - Угадай Число - \n"
	fmt.Print(text)
	file.WriteString(text)
	text = "Попытайтесь угадать задуманное число из диапазона от 1 до 100. \n"
	fmt.Print(text)
	file.WriteString(text)
	text = "Программа будет подсказывать введенное число меньше или больше задуманого. \n"
	fmt.Print(text)
	file.WriteString(text)

	for answerNumber != conceiveNumber {
		text = "Введите число: \n"
		fmt.Print(text)
		file.WriteString(text)
		fmt.Scan(&answerNumber)
		text = "%d \n"
		file.WriteString(fmt.Sprintf(text, answerNumber))
		if answerNumber < 1 || answerNumber > 100 {
			text = "Число должно быть из диапазона от 1 до 100! \n"
			fmt.Print(text)
			file.WriteString(text)
			continue
		} else if answerNumber < conceiveNumber {
			text = "Задуманное число больше. Попробуйте еше раз! \n"
			fmt.Print(text)
			file.WriteString(text)
		} else if answerNumber > conceiveNumber {
			text = "Задуманное число меньше. Попробуйте еше раз! \n"
			fmt.Print(text)
			file.WriteString(text)
		}
	}
	text = "Вы угадали! Это число: %d \n"
	fmt.Printf(text, conceiveNumber)
	file.WriteString(fmt.Sprintf(text, conceiveNumber))
}
