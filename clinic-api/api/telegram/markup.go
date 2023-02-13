package telegram

import (
	tgBot "github.com/Syfaro/telegram-bot-api"
	"time"
)

const (
	nextPage     = ">>"
	previousPage = "<<"
)

func yearsPaginationMarkup(startDate time.Time) tgBot.InlineKeyboardMarkup {
	return generatePaginationMarkup([][]string{
		{startDate.AddDate(-8, 0, 0).Format("2006"), startDate.AddDate(-7, 0, 0).Format("2006"), startDate.AddDate(-6, 0, 0).Format("2006")},
		{startDate.AddDate(-5, 0, 0).Format("2006"), startDate.AddDate(-4, 0, 0).Format("2006"), startDate.AddDate(-3, 0, 0).Format("2006")},
		{startDate.AddDate(-2, 0, 0).Format("2006"), startDate.AddDate(-1, 0, 0).Format("2006"), startDate.Format("2006")},
	})
}

func generatePaginationMarkup(data [][]string) tgBot.InlineKeyboardMarkup {
	mkup := tgBot.NewInlineKeyboardMarkup()
	for x := range data {
		var a []tgBot.InlineKeyboardButton
		for y := range data[x] {
			b := tgBot.NewInlineKeyboardButtonData(data[x][y], data[x][y])
			a = append(a, b)
		}
		mkup.InlineKeyboard = append(mkup.InlineKeyboard, a)
	}
	mkup.InlineKeyboard = append(mkup.InlineKeyboard, []tgBot.InlineKeyboardButton{tgBot.NewInlineKeyboardButtonData("<<", previousPage), tgBot.NewInlineKeyboardButtonData(">>", nextPage)})

	return mkup
}

func generateMarkupByDays(daysCount int) tgBot.InlineKeyboardMarkup {
	mkup := tgBot.NewInlineKeyboardMarkup(
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData("1", "1"),
			tgBot.NewInlineKeyboardButtonData("2", "2"),
			tgBot.NewInlineKeyboardButtonData("3", "3"),
			tgBot.NewInlineKeyboardButtonData("4", "4"),
			tgBot.NewInlineKeyboardButtonData("5", "5"),
			tgBot.NewInlineKeyboardButtonData("6", "6"),
			tgBot.NewInlineKeyboardButtonData("7", "1"),
		), tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData("8", "8"),
			tgBot.NewInlineKeyboardButtonData("9", "9"),
			tgBot.NewInlineKeyboardButtonData("10", "10"),
			tgBot.NewInlineKeyboardButtonData("11", "11"),
			tgBot.NewInlineKeyboardButtonData("12", "12"),
			tgBot.NewInlineKeyboardButtonData("13", "13"),
			tgBot.NewInlineKeyboardButtonData("14", "14"),
		), tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData("15", "15"),
			tgBot.NewInlineKeyboardButtonData("16", "16"),
			tgBot.NewInlineKeyboardButtonData("17", "17"),
			tgBot.NewInlineKeyboardButtonData("18", "18"),
			tgBot.NewInlineKeyboardButtonData("19", "19"),
			tgBot.NewInlineKeyboardButtonData("20", "20"),
			tgBot.NewInlineKeyboardButtonData("21", "21"),
		), tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData("22", "22"),
			tgBot.NewInlineKeyboardButtonData("23", "23"),
			tgBot.NewInlineKeyboardButtonData("24", "24"),
			tgBot.NewInlineKeyboardButtonData("25", "25"),
			tgBot.NewInlineKeyboardButtonData("26", "26"),
			tgBot.NewInlineKeyboardButtonData("27", "27"),
			tgBot.NewInlineKeyboardButtonData("28", "28"),
		))

	if daysCount == 29 {
		mkup.InlineKeyboard = append(mkup.InlineKeyboard, []tgBot.InlineKeyboardButton{
			tgBot.NewInlineKeyboardButtonData("29", "29"),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " ")})
	} else if daysCount == 30 {
		mkup.InlineKeyboard = append(mkup.InlineKeyboard, []tgBot.InlineKeyboardButton{
			tgBot.NewInlineKeyboardButtonData("29", "29"),
			tgBot.NewInlineKeyboardButtonData("30", "30"),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " ")})
	} else if daysCount == 31 {
		mkup.InlineKeyboard = append(mkup.InlineKeyboard, []tgBot.InlineKeyboardButton{
			tgBot.NewInlineKeyboardButtonData("29", "29"),
			tgBot.NewInlineKeyboardButtonData("30", "30"),
			tgBot.NewInlineKeyboardButtonData("31", "31"),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " "),
			tgBot.NewInlineKeyboardButtonData(" ", " ")})
	}

	return mkup
}
