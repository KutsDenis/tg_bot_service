package alert

import (
	"tg_bot_template/internal/cfg"
	"tg_bot_template/internal/messenger"
)

type Sender interface {
	ToAdmin()
	ToGroup()
}

type Alert struct {
	Text     string
	Markdown bool
}

func (a Alert) ToAdmin() {
	message := messenger.Message{
		ID:       cfg.Get.Bot.Admin,
		Text:     a.Text,
		Markdown: a.Markdown,
	}
	go message.Send()
}

func (a Alert) ToGroup() {
	message := messenger.Message{
		ID:       cfg.Get.Bot.Group,
		Text:     a.Text,
		Markdown: a.Markdown,
	}
	go message.Send()
}
