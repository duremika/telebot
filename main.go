package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"time"
)

var token = "1488492813:AAFWcpzyTh4aOeP5anW--ShkxHD3NaJLhMI"

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120

	updates, err := bot.GetUpdatesChan(u)

	CalculateKeyboard()
	userEditDate := make(map[int64]int)
	for update := range updates {
		if update.CallbackQuery != nil {
			data := update.CallbackQuery.Data
			if data == "b" {
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Чем могу помочь?")
				msg.ReplyMarkup = viewButton
				bot.Send(msg)
			} else if data == "v" {
				bot.DeleteMessage(tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID-2))
				warning := ""
				aWeekAgo := time.Now().AddDate(0, 0, 7)
				for _, x := range equips {
					if x.Date.Before(aWeekAgo) {
						warning += "‼️" + x.Name + "‼️следующая дата испытаний меньше чем через неделю‼️\n"
					}
				}
				/*
					bot.Send(tgbotapi.NewEditMessageTextAndMarkup(
						update.CallbackQuery.Message.Chat.ID,
						update.CallbackQuery.Message.MessageID,
						warning+"\nВыберите средство защиты для изменения даты",
						equipKeyboard,
					))
				*/
				msg := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					warning+"\nВыберите средство защиты для изменения даты",
				)
				msg.ReplyMarkup = &equipKeyboard
				bot.Send(msg)
			} else if i, e := strconv.Atoi(data); e == nil {
				userEditDate[update.CallbackQuery.Message.Chat.ID] = i
				/*
					bot.Send(tgbotapi.NewEditMessageTextAndMarkup(
						update.CallbackQuery.Message.Chat.ID,
						update.CallbackQuery.Message.MessageID,
						"Выберите день",
						dayKeyboard,
					))
				*/
				msg := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"Выберите день",
				)
				msg.ReplyMarkup = &dayKeyboard
				bot.Send(msg)
			} else if string(data[0]) == "d" {
				d, _ := strconv.ParseInt(data[1:], 10, 16)
				date := equips[userEditDate[update.CallbackQuery.Message.Chat.ID]].Date
				equips[userEditDate[update.CallbackQuery.Message.Chat.ID]].Date =
					time.Date(date.Year(), date.Month(), int(d), 0, 0, 0, 0, time.Local)
				/*
					bot.Send(tgbotapi.NewEditMessageTextAndMarkup(
						update.CallbackQuery.Message.Chat.ID,
						update.CallbackQuery.Message.MessageID,
						"Выберите месяц",
						mounthKeyboard,
					))
				*/
				msg := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"Выберите месяц",
				)
				msg.ReplyMarkup = &mounthKeyboard
				bot.Send(msg)
			} else if string(data[0]) == "m" {
				m, _ := strconv.ParseInt(data[1:], 10, 16)
				date := equips[userEditDate[update.CallbackQuery.Message.Chat.ID]].Date
				equips[userEditDate[update.CallbackQuery.Message.Chat.ID]].Date =
					time.Date(date.Year(), time.Month(m), date.Day(), 0, 0, 0, 0, time.Local)
				/*
					bot.Send(tgbotapi.NewEditMessageTextAndMarkup(
						update.CallbackQuery.Message.Chat.ID,
						update.CallbackQuery.Message.MessageID,
						"Выберите год",
						yearKeyboard,
					))
				*/
				msg := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"Выберите год",
				)
				msg.ReplyMarkup = &yearKeyboard
				bot.Send(msg)
			} else if string(data[0]) == "y" {
				equip := equips[userEditDate[update.CallbackQuery.Message.Chat.ID]]
				date := equip.Date
				year, _ := strconv.ParseInt(string(data[1]), 10, 16)
				date = time.Date(2020+int(year), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
				equip.Date = date

				Update(&equip)
				RecalculateKeyboard()
				/*
					bot.Send(tgbotapi.NewEditMessageTextAndMarkup(
						update.CallbackQuery.Message.Chat.ID,
						update.CallbackQuery.Message.MessageID,
						"Принято",
						viewButton,
					))
				*/
				msg := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					"Принято",
				)
				msg.ReplyMarkup = &viewButton
				bot.Send(msg)
			}
			continue
		}
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Чем могу помочь?")
			msg.ReplyMarkup = viewButton
			bot.Send(msg)
		}
	}
}
