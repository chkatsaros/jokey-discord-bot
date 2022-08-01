package main

import (
	"fmt"
	"jokey-discord-bot/bot"
	"jokey-discord-bot/config"
	"net/http"
	"os"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
