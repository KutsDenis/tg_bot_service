package commands

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg_bot_template/internal/messenger"
)

type Commander interface {
	CallCommand(command Command)
}

type Command struct {
	CMD  string
	User *botapi.User
	Bot  *botapi.BotAPI
}

func (c Command) CallCommand() {
	msg := messenger.Message{
		ID:       c.User.ID,
		Text:     "",
		Markdown: false,
	}

	switch c.CMD {
	case "help":
		msg.Text = "Что сюда написать?"
	case "start":
		msg.Text = "https://memepedia.ru/wp-content/uploads/2017/05/%D1%8F-%D1%81%D0%BA%D0%B0%D0%B7%D0%B0%D0%BB%D0%B0-%D1%81%D1%82%D0%B0%D1%80%D1%82%D1%83%D0%B5%D0%BC.jpg"
	default:
		msg.Text = "Неизвестная команда"
	}

	if msg.Text != "" {
		msg.Send()
	}
}
