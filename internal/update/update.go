package update

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg_bot_template/cmd/main/config"
	"tg_bot_template/internal/commands"
)

type Updater interface {
	Update()
}

type Update struct {
}

func (u Update) Update() {
	var init config.Initializer
	var command commands.Commander

	init = config.Config{}
	command = commands.Command{}

	bot := init.Init()
	cmd := command.Command

	upd := botapi.NewUpdate(0)
	upd.Timeout = 60

	updates := bot.GetUpdatesChan(upd)

	for update := range updates {
		usr := update.SentFrom()

		/* Перехват колбэков */
		if update.CallbackQuery != nil {
			// todo
		}

		/* Игнорирование любых обновлений кроме сообщений */
		if update.Message == nil {
			continue
		}

		/* Игнорирование всего кроме команд */
		if !update.Message.IsCommand() {
			continue
		}

		// Команды
		go cmd(update.Message.Command(), usr)
	}
}
