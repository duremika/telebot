package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"time"
)

func HelloHandler(id int64) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(
		id,
		"Чем могу помочь?",
	)
	msg.ReplyMarkup = viewButton
	return msg
}

func Handle(data string, id int64, mId int) tgbotapi.Chattable {
	var txt string
	var markup tgbotapi.InlineKeyboardMarkup

	if i, err := strconv.Atoi(data); err == nil {
		equip = equips[i]
		txt = "Выберите день"
		markup = dayKeyboard
	} else if data == "v" {
		go bot.DeleteMessage(tgbotapi.NewDeleteMessage(id, mId-2))
		warning := ""
		aWeekAgo := time.Now().AddDate(0, 0, 7)
		for _, x := range equips {
			if x.Date.Before(aWeekAgo) {
				warning += "‼️" + x.Name + "‼️следующая дата испытаний меньше чем через неделю‼️\n"
			}
		}
		txt = warning + "\nВыберите средство защиты для изменения даты"
		markup = equipKeyboard
	} else {
		date := equip.Date
		num, err := strconv.Atoi(data)
		if err != nil {
			return nil
		}

		switch string(data[0]) {
		case "d":
			date = time.Date(date.Year(), date.Month(), num, 0, 0, 0, 0, time.Local)
			txt = "Выберите месяц"
			markup = mounthKeyboard
		case "m":
			date = time.Date(date.Year(), time.Month(num), date.Day(), 0, 0, 0, 0, time.Local)
			txt = "Выберите год"
			markup = yearKeyboard
		case "y":
			date = time.Date(num, date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
			RecalculateKeyboard(equip)
			txt = "Принято"
			markup = viewButton
		}
	}

	msg := tgbotapi.NewEditMessageText(id, mId, txt)
	msg.ReplyMarkup = &markup
	return msg
}
