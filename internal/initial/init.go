package initial

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
	"tg_bot_template/internal/cfg"
)

var onceInit sync.Once
var Bot *botapi.BotAPI

// Init инициализирует бота.
func Init() {
	onceInit.Do(func() {
		var err error

		Bot, err = botapi.NewBotAPI(cfg.Get.Bot.Token)
		if err != nil {
			log.Panic(err)
		}

		Bot.Debug = true
	})
}
