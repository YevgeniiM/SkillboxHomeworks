package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

//goland:noinspection ALL
func main() {

	file, err := os.Create("Message.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("Программа выводит сообщения в формате:\n№ строки, дата, сообщение \n")

	in := bufio.NewReader(os.Stdin)
	//var message string
	//var codLit int

	for j := 1; ; j++ {
		//message = ""
		//fmt.Println("Введите сообщение:")
		//for { // собитаем строку посимвольно через unicod до символа переноса строки
		//	_, _ = fmt.Scanf("%c", &codLit)
		//	if string(codLit) == "\n" {
		//		break
		//	}
		//	message += string(codLit)
		//	//fmt.Printf("%s", message)
		//}
		//if message == "выход" {
		//	file.WriteString("Вышли из программы \n")
		//	return
		//} else {
		//	dt := time.Now()
		//	file.WriteString(fmt.Sprintf("№ %-3d %s    %s \n", j, dt.Format("01-02-2006 15:04:05"), message))

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
}
