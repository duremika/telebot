package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
)

var (
	bot   tgbotapi.BotAPI
	equip Equip
)

func main() {
	port := os.Getenv("PORT")
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()
	bot, _ := tgbotapi.NewBotAPI(token)
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	if _, err := bot.SetWebhook(tgbotapi.NewWebhook(webhook)); err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")

	CalculateKeyboard()
	for update := range updates {

		var id int64
		if update.CallbackQuery == nil {
			id = update.Message.Chat.ID
			bot.Send(HelloHandler(id))
			continue
		}

		data := update.CallbackQuery.Data

		if update.CallbackQuery != nil {

			id = update.CallbackQuery.Message.Chat.ID
			if data == "b" {
				bot.Send(HelloHandler(id))
				continue
			}
		}

		mId := update.CallbackQuery.Message.MessageID
		msg := Handle(data, id, mId)
		if _, err := bot.Send(msg); err != nil {
			log.Fatal(err)
		}
	}
}
