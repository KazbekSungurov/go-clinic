package telegram

import (
	"clinic-api/models"
	"clinic-api/service"
	"fmt"
	tgBot "github.com/Syfaro/telegram-bot-api"
	"strings"
)

var (
	createMarkup = tgBot.NewInlineKeyboardMarkup(
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData("Подтвердить", "Подтвердить"),
			tgBot.NewInlineKeyboardButtonData("Изменить", "Изменить"),
		))
)

type patient struct {
	base
	toDo  func(u *tgBot.Update)
	model *models.Patient
}

func initCreatePatient(id int64, api *tgBot.BotAPI, service *service.Service) *patient {
	var c patient
	c.Service = service
	c.ID, c.api = id, api
	c.timerSet(sessionTime)
	c.create()
	return &c
}

func (c *patient) exec(u *tgBot.Update) {
	c.toDo(u)
}

func (c *patient) create() {
	c.model = &models.Patient{}
	c.toDo = func(u *tgBot.Update) {
		switch c.lvl {
		case 0:
			msg := tgBot.NewMessage(c.ID, "Введите Ваше имя:")
			_, err := c.api.Send(msg)
			if err != nil {
				return
			}
			c.lvl++
		case 1:
			c.model.FirstName = strings.Title(strings.ToLower(u.Message.Text))
			msg := tgBot.NewMessage(c.ID, "Введите Вашу фамилию:")
			c.api.Send(msg)
			c.lvl++
		case 2:
			c.model.LastName = strings.Title(strings.ToLower(u.Message.Text))
			msg := tgBot.NewMessage(c.ID, "Введите Ваше Отчество:")
			c.api.Send(msg)
			c.lvl++
		case 3:
			c.model.MiddleName = strings.Title(strings.ToLower(u.Message.Text))
			//msg := tgBot.NewMessage(int64(u.Message.From.ID), "Введите Вашу дату рождения в формате дд-мм-гггг (10-10-2000):")
			//c.api.Send(msg)
			c.lvl++
			cldr := newCalendar(c, c.api)
			cldr.date = &c.model.BirthDate
			cldr.exec(u)
		case 4:
			msg := tgBot.NewMessage(c.ID, "Введите Ваш адрес:")
			c.api.Send(msg)
			c.lvl++
		case 5:
			c.model.Address = u.Message.Text
			msg := tgBot.NewMessage(c.ID, "Введите Ваш номер телефона:")
			c.api.Send(msg)
			c.lvl++
		case 6:
			c.model.Phone = u.Message.Text
			msg := tgBot.NewMessage(c.ID, fmt.Sprintf(
				`Проверьте введенные данные:
Имя: %s
Фамилия: %s
Отчество: %s
Дата рождения: %s
Адрес: %s
Телефон: %s`,
				c.model.FirstName, c.model.LastName, c.model.MiddleName, c.model.BirthDate.Format("02-01-2006"), c.model.Address, c.model.Phone))
			msg.ReplyMarkup = createMarkup
			c.api.Send(msg)
			c.lvl++
		case 7:
			if u.Message != nil {
				msg := tgBot.NewMessage(c.ID, "Пожалуйста, выберите один из вариантов")
				c.api.Send(msg)
			} else if u.CallbackQuery != nil {
				switch u.CallbackQuery.Data {
				case "Подтвердить":
					fmt.Println(c.model)
					if _, err := c.Patient.Create(c.model); err != nil {
						return
					}
					endSession(c.ID)
					msg := tgBot.NewMessage(c.ID, "Успешно создано!")
					c.api.Send(msg)
				case "Изменить":
					msg := tgBot.NewMessage(c.ID, "Введите Ваше имя:")
					c.api.Send(msg)
					c.lvl = 1
				}
			}
		}
	}
}
