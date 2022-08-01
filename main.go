package main

import (
	"fmt"
	"jokey-discord-bot/bot"
	"jokey-discord-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
}
