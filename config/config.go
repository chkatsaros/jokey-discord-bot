package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json : "Token"`
	BotPrefix string `json : "BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	/* ioutil package's ReadFile method which we read config.json and return it's value we will then
	store it in file variable and if an error ocurrs it will be stored in err . */
	file, err := ioutil.ReadFile("./config.json")

	// Handling error and printing it using fmt package's Println function and returning it .
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// We are here printing value of file variable by explicitly converting it to string .

	fmt.Println(string(file))
	/* Here we performing a simple task by copying value of file into config variable which we have declared above,
	and if there any error we are storing it in err . Unmarshal takes second arguments reference remember it. */
	err = json.Unmarshal(file, &config)

	// Handling error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// After storing value in config variable we will access it and storing it in our declared variables.
	Token = config.Token
	BotPrefix = config.BotPrefix

	// If there isn't any error we will return nil.
	return nil

}
