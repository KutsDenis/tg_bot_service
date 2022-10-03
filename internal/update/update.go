package update

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg_bot_template/internal/commands"
)

// Update Обновление событий бота
func Update(bot *botapi.BotAPI) {
	upd := botapi.NewUpdate(0)
	upd.Timeout = 60

	updates := bot.GetUpdatesChan(upd)

	for update := range updates {
		usr := update.SentFrom()

		// Перехват колбэков
		if update.CallbackQuery != nil {
			// todo
		}

		// Игнорирование любых обновлений кроме сообщений
		if update.Message == nil {
			continue
		}

		// Игнорирование всего кроме команд
		if !update.Message.IsCommand() {
			continue
		}

		// Команды
		cmd := commands.Command{
			CMD:  update.Message.Command(),
			User: usr,
			Bot:  bot,
		}

		cmd.CallCommand()
	}
}
