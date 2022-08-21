package messenger

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Send(id int64, text string, bot *botapi.BotAPI) {
	msg := botapi.NewMessage(id, text)

	if _, err := bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}
