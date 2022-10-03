package messenger

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tg_bot_template/internal/initial"
)

type Messenger interface {
	Send()
}

type Message struct {
	ID       int64
	Text     string
	Markdown bool
}

// Send Отправка сообщений.
func (m Message) Send() {
	bot := initial.Bot

	msg := botapi.NewMessage(m.ID, m.Text)

	if m.Markdown {
		msg.ParseMode = "markdown"
	}

	if _, err := bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}
