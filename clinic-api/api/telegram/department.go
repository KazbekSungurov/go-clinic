package telegram

import (
	"clinic-api/models"
	"clinic-api/service"
	"fmt"
	tgBot "github.com/Syfaro/telegram-bot-api"
	"time"
)

type department struct {
	base
	toDo     func(u *tgBot.Update)
	model    *models.Department
	from, to time.Time
}

func initDepartment(id int64, api *tgBot.BotAPI, service *service.Service) *department {
	var c department
	c.Service = service
	c.ID, c.api = id, api
	c.timerSet(sessionTime)
	c.create()
	return &c
}

func (c *department) exec(u *tgBot.Update) {
	c.toDo(u)
}

func (c *department) create() {
	c.model = &models.Department{}
	c.toDo = func(u *tgBot.Update) {
		switch c.lvl {
		case 0:
			msg := tgBot.NewMessage(c.ID, "Выберите диапазон дат")
			_, err := c.api.Send(msg)
			if err != nil {
				return
			}
			c.lvl++
			cldr := newCalendar(c, c.api)
			cldr.date = &c.from
			cldr.exec(u)
		case 1:
			c.lvl++
			cldr := newCalendar(c, c.api)
			cldr.date = &c.to
			cldr.exec(u)
		case 2:
			c.model.ScheduleOperation.WeekDays = append(c.model.ScheduleOperation.WeekDays, time.Monday, time.Tuesday, time.Wednesday, time.Thursday)
			err := c.Service.ScheduleOperationBuild(c.model, c.from, c.to)
			if err != nil {
				fmt.Println(err)
			}
			var s string
			for i := range c.model.ScheduleOperation.Appointments {
				s += fmt.Sprintf("%v\n", c.model.ScheduleOperation.Appointments[i])
			}
			msg := tgBot.NewMessage(c.ID, s)
			_, err = c.api.Send(msg)
			if err != nil {
				return
			}
			endSession(c.ID)
		}
	}
}
