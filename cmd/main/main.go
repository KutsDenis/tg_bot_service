package main

import "tg_bot_template/internal/update"

func main() {
	var u update.Updater

	u = update.Update{}
	u.Update()
}
