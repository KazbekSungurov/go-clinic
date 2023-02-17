package telegram

import (
	"clinic-api/service"
	tgBot "github.com/Syfaro/telegram-bot-api"
)

const (
	botToken string = "5818415012:AAHOO2abbEn5BnDuJhXOkBe65SZ28PVAVxs"
)

var (
	mainMarkup = tgBot.NewInlineKeyboardMarkup(
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData("Создать клиента", "Создать клиента"),
		),
		tgBot.NewInlineKeyboardRow(
			tgBot.NewInlineKeyboardButtonData("Ничего не надо!", "Ничего не надо!"),
		),
	)
)

type Bot struct {
	api *tgBot.BotAPI
	*sessions
}

func New(s *service.Service) (*Bot, error) {
	api, err := tgBot.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:      api,
		sessions: newSessions(s),
	}, nil
}

func (b *Bot) Init() error {
	u := tgBot.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.api.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message != nil {
			s := b.sessions.check(update.Message.Chat.ID)
			if s != nil {
				s.exec(&update)
			} else {
				b.defaultMsg(&update)
			}
		} else if update.CallbackQuery != nil {
			s := b.sessions.check(int64(update.CallbackQuery.From.ID))
			if s != nil {
				s.exec(&update)
			} else {
				b.defaultCallBack(&update)
			}
		}
	}
	return nil
}

func (b *Bot) defaultCallBack(update *tgBot.Update) {
	//callback := tgBot.NewCallback(update.CallbackQuery.ID, "✓")
	//b.bot.AnswerCallbackQuery(callback)

	switch update.CallbackQuery.Data {
	case "Создать клиента":
		s := b.initDepartment(int64(update.CallbackQuery.From.ID), b.api)
		s.exec(update)
	default:
	}
}

func (b *Bot) defaultMsg(update *tgBot.Update) {
	msg := tgBot.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyMarkup = mainMarkup
	b.api.Send(msg)
}
