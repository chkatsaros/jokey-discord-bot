package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token     string
	BotPrefix string
)

func ReadConfig() error {
	fmt.Println("Reading config file...")

	Token = goDotEnvVariable("token")
	BotPrefix = goDotEnvVariable("botPrefix")

	fmt.Println("==> Token: ", Token)
	fmt.Println("==> BotPrefix: ", BotPrefix)

	// If there isn't any error we will return nil.
	return nil

}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
