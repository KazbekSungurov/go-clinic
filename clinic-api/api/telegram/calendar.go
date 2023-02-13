package telegram

import (
	"fmt"
	tgBot "github.com/Syfaro/telegram-bot-api"
	"strconv"
	"time"
)

type calendar struct {
	base
	parentSession    session
	date             *time.Time
	year, month, day string
}

var (
	t          = time.Now()
	yearMarkUp = tgBot.NewInlineKeyboardMarkup(
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-18, 0, 0).Format("2006"), t.AddDate(-18, 0, 0).Format("2006")),
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-19, 0, 0).Format("2006"), t.AddDate(-19, 0, 0).Format("2006")),
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-20, 0, 0).Format("2006"), t.AddDate(-20, 0, 0).Format("2006")),
		),
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-22, 0, 0).Format("2006"), t.AddDate(-22, 0, 0).Format("2006")),
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-23, 0, 0).Format("2006"), t.AddDate(-23, 0, 0).Format("2006")),
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-24, 0, 0).Format("2006"), t.AddDate(-24, 0, 0).Format("2006")),
		),
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-26, 0, 0).Format("2006"), t.AddDate(-26, 0, 0).Format("2006")),
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-27, 0, 0).Format("2006"), t.AddDate(-27, 0, 0).Format("2006")),
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-28, 0, 0).Format("2006"), t.AddDate(-28, 0, 0).Format("2006")),
		),
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-30, 0, 0).Format("2006"), t.AddDate(-30, 0, 0).Format("2006")),
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-31, 0, 0).Format("2006"), t.AddDate(-31, 0, 0).Format("2006")),
			tgBot.NewInlineKeyboardButtonData(t.AddDate(-32, 0, 0).Format("2006"), t.AddDate(-32, 0, 0).Format("2006")),
		))

	monthMarkUp = tgBot.NewInlineKeyboardMarkup(
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData("Январь", "01"),
			tgBot.NewInlineKeyboardButtonData("Февраль", "02"),
			tgBot.NewInlineKeyboardButtonData("Март", "03"),
		),
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData("Апрель", "04"),
			tgBot.NewInlineKeyboardButtonData("Май", "05"),
			tgBot.NewInlineKeyboardButtonData("Июнь", "06"),
		),
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData("Июль", "07"),
			tgBot.NewInlineKeyboardButtonData("Август", "08"),
			tgBot.NewInlineKeyboardButtonData("Сентябрь", "09"),
		),
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData("Октябрь", "10"),
			tgBot.NewInlineKeyboardButtonData("Ноябрь", "11"),
			tgBot.NewInlineKeyboardButtonData("Декабрь", "12"),
		))
)

func newCalendar(s session, api *tgBot.BotAPI) *calendar {
	c := calendar{
		base: base{
			ID:  s.id(),
			api: api,
		},
		parentSession: s,
	}

	c.timerSet(sessionTime)
	switchSessions(&c)

	return &c
}

func (c *calendar) exec(u *tgBot.Update) {
	switch c.lvl {
	case 0:
		*c.date = time.Now()
		msg := tgBot.NewMessage(c.ID, "Выберите дату:")
		msg.ReplyMarkup = yearsPaginationMarkup(c.date.AddDate(-18, 0, 0))
		c.api.Send(msg)
		c.lvl++
	case 1:
		if u.CallbackQuery != nil {
			if u.CallbackQuery.Data == previousPage {
				*c.date = c.date.AddDate(-9, 0, 0)
				msg := tgBot.NewEditMessageReplyMarkup(c.ID, u.CallbackQuery.Message.MessageID, yearsPaginationMarkup(*c.date))
				_, err := c.api.Send(msg)
				if err != nil {
					fmt.Println(err)
				}
			} else if u.CallbackQuery.Data == nextPage {
				*c.date = c.date.AddDate(9, 0, 0)
				msg := tgBot.NewEditMessageReplyMarkup(c.ID, u.CallbackQuery.Message.MessageID, yearsPaginationMarkup(*c.date))
				_, err := c.api.Send(msg)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				*c.date, _ = time.Parse("02.01.2006", fmt.Sprintf("01.01.%s", u.CallbackQuery.Data))
				c.year = u.CallbackQuery.Data
				msg := tgBot.NewEditMessageReplyMarkup(c.ID, u.CallbackQuery.Message.MessageID, monthMarkUp)
				_, err := c.api.Send(msg)
				if err != nil {
					fmt.Println(err)
				}
				c.lvl++
			}
		}
	case 2:
		t, _ := strconv.Atoi(u.CallbackQuery.Data)
		*c.date = c.date.AddDate(0, t-1, 0)
		msg := tgBot.NewEditMessageReplyMarkup(c.ID, u.CallbackQuery.Message.MessageID, generateMarkupByDays(daysInMonth(*c.date)))
		_, err := c.api.Send(msg)
		if err != nil {
			fmt.Println(err)
		}
		c.lvl++
	case 3:
		t, _ := strconv.Atoi(u.CallbackQuery.Data)
		*c.date = c.date.AddDate(0, 0, t-1)
		msg := tgBot.NewMessage(c.ID, fmt.Sprintf(c.date.Format("02.01.2006")))
		_, err := c.api.Send(msg)
		if err != nil {
			fmt.Println(err)
		}
		switchSessions(c.parentSession)
		c.parentSession.exec(u)
	}
}

func daysInMonth(t time.Time) int {
	y, m, _ := t.Date()
	return time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
