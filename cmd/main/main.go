package main

import (
	"tg_bot_template/internal/cfg"
	"tg_bot_template/internal/initial"
	"tg_bot_template/internal/update"
)

func main() {
	cfg.Load()
	initial.Init()
	bot := initial.Bot

	update.Update(bot)
}
