package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

var (
	equipsDB   []Equip
	viewButton = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Посмотреть", "v"),
		),
	)

	cancelButton = tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Назад", "b"),
	)
	equips []Equip
	layout = "02 01 06"

	equipKeyboard = tgbotapi.InlineKeyboardMarkup{[][]tgbotapi.InlineKeyboardButton{}}

	dayKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "d1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "d2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "d3"),
			tgbotapi.NewInlineKeyboardButtonData("4", "d4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "d5"),
			tgbotapi.NewInlineKeyboardButtonData("6", "d6"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("7", "d7"),
			tgbotapi.NewInlineKeyboardButtonData("8", "d8"),
			tgbotapi.NewInlineKeyboardButtonData("9", "d9"),
			tgbotapi.NewInlineKeyboardButtonData("10", "d10"),
			tgbotapi.NewInlineKeyboardButtonData("11", "d11"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("12", "d12"),
			tgbotapi.NewInlineKeyboardButtonData("13", "d13"),
			tgbotapi.NewInlineKeyboardButtonData("14", "d14"),
			tgbotapi.NewInlineKeyboardButtonData("15", "d15"),
			tgbotapi.NewInlineKeyboardButtonData("16", "d16"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("17", "d17"),
			tgbotapi.NewInlineKeyboardButtonData("18", "d18"),
			tgbotapi.NewInlineKeyboardButtonData("19", "d19"),
			tgbotapi.NewInlineKeyboardButtonData("20", "d20"),
			tgbotapi.NewInlineKeyboardButtonData("21", "d21"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("22", "d22"),
			tgbotapi.NewInlineKeyboardButtonData("23", "d23"),
			tgbotapi.NewInlineKeyboardButtonData("24", "d24"),
			tgbotapi.NewInlineKeyboardButtonData("25", "d25"),
			tgbotapi.NewInlineKeyboardButtonData("26", "d26"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("27", "d27"),
			tgbotapi.NewInlineKeyboardButtonData("28", "d28"),
			tgbotapi.NewInlineKeyboardButtonData("29", "d29"),
			tgbotapi.NewInlineKeyboardButtonData("30", "d30"),
			tgbotapi.NewInlineKeyboardButtonData("31", "d31"),
		),
	)

	mounthKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Январь", "m1"),
			tgbotapi.NewInlineKeyboardButtonData("Февраль", "m2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Март", "m3"),
			tgbotapi.NewInlineKeyboardButtonData("Апрель", "m4"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Май", "m5"),
			tgbotapi.NewInlineKeyboardButtonData("Июнь", "m6"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Июль", "m7"),
			tgbotapi.NewInlineKeyboardButtonData("Август", "m8"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Сентябрь", "m9"),
			tgbotapi.NewInlineKeyboardButtonData("Октябрь", "m10"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ноябрь", "m11"),
			tgbotapi.NewInlineKeyboardButtonData("Декабрь", "m12"),
		),
	)
	yearKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("2021", "y1"),
			tgbotapi.NewInlineKeyboardButtonData("2022", "y2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("2023", "y3"),
			tgbotapi.NewInlineKeyboardButtonData("2024", "y4"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("2025", "y5"),
			tgbotapi.NewInlineKeyboardButtonData("2026", "y6"),
		),
	)
)

//func RecalculateKeyboard() {
//	equipsDB, er := FindAll()
//	if er != nil {
//		println(er.Error())
//	}
//	equips = equipsDB
//	for i, _ := range equipKeyboard.InlineKeyboard {
//		equipKeyboard.InlineKeyboard[i][0] =
//			tgbotapi.NewInlineKeyboardButtonData(equips[i].Name+" : "+equips[i].Date.Format(layout), strconv.Itoa(i))
//	}
//}

func RecalculateKeyboard(equip Equip) {
	go func() {
		equipsDB, _ = FindAll()
		equips = equipsDB
	}()
	go func() {
		Update(&equip)
	}()
	for i, _ := range equipKeyboard.InlineKeyboard {
		equipKeyboard.InlineKeyboard[i][0] =
			tgbotapi.NewInlineKeyboardButtonData(equips[i].Name+" : "+equips[i].Date.Format(layout), strconv.Itoa(i))
	}
}

func CalculateKeyboard() {
	equipsDB, _ = FindAll()
	equips = equipsDB
	for i, _ := range equips {
		equipKeyboard.InlineKeyboard = append(equipKeyboard.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(equips[i].Name+" : "+equips[i].Date.Format(layout), strconv.Itoa(i)),
			),
		)
	}

	dayKeyboard.InlineKeyboard = append(dayKeyboard.InlineKeyboard, cancelButton)
	mounthKeyboard.InlineKeyboard = append(mounthKeyboard.InlineKeyboard, cancelButton)
	yearKeyboard.InlineKeyboard = append(yearKeyboard.InlineKeyboard, cancelButton)
}
