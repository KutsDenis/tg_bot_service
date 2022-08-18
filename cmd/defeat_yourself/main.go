package main

import "defeat-yourself/internal/update"

func main() {
	var u update.Updater

	u = update.Update{}
	u.Update()
}
