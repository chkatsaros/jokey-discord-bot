package config

import (
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string
)

func ReadConfig() error {
	fmt.Println("Reading config file...")

	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")

	fmt.Println("==> Token: ", Token)
	fmt.Println("==> BotPrefix: ", BotPrefix)

	// If there isn't any error we will return nil.
	return nil

}
