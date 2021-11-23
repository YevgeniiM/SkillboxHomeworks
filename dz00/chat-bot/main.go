package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type familyBudget map[string]int64

var bd = map[int]familyBudget{}

func budget(userId int) {
	message := ""
	f, err := os.Create("familyBudget.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for k, v := range bd[userId] {
		message += fmt.Sprintf("%s %v грн. \n", k, v)

	}
	f.WriteString(message)
	//bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, message1))

}

func main() {
	bot, err := tgbotapi.NewBotAPI("2064436754:AAEwxBra4yqxHpEJPtMYNq76UNGP92ET4fA")
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true
	//
	//log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		var message1 string
		var command []string
		var sum int64
		var userId int
		command = strings.Split(update.Message.Text, " ")
		userId = update.Message.From.ID

		switch command[0] {
		case "категорії":
			message1 = ""
			for k, _ := range bd[userId] {
				//fmt.Println(k, v)
				message1 += fmt.Sprintf("%s. \n", k)
			}
			if message1 == "" {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "поки нема жодної категорії"))
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, message1))
			}

			//if len(command) != 3 {
			//	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, message1))
			//}
		case "+":
			if len(command) != 3 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
			}
			if _, ok := bd[userId]; !ok {
				bd[userId] = make(familyBudget)
			}
			sum, err = strconv.ParseInt(command[2], 10, 64)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "не верный ввод! добавляемая сумма должна быть целым числом"))
			} else {
				bd[userId][command[1]] += sum
				budget(userId)
			}

		case "видалити":
			if len(command) != 2 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Для видалення категорії пишемо ʼвидалитиʼ и через пробіл назву категорії - їх можливо побачити нарисавши 'бюджетʼ  "))
				continue
			}
			delete(bd[userId], command[1])
		case "бюджет":
			message1 = ""
			for k, v := range bd[userId] {
				//fmt.Println(k, v)
				message1 += fmt.Sprintf("%s %v грн. \n", k, v)
			}
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, message1))
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		//bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
	}
}
