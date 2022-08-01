package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jokey-discord-bot/config"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

// var jokeAPI string = "https://v2.jokeapi.dev/joke/"
var jokeAPI string = "https://api.jokes.one"

type JokeApiResponse struct {
	Success struct {
		Total int `json:"total"`
	}
	Contents struct {
		Jokes []struct {
			Category    string `json:"category"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Background  string `json:"background"`
			Date        string `json:"date"`
			Joke        struct {
				Title  string `json:"title"`
				Length string `json:"length"`
				Clean  string `json:"clean"`
				Racial string `json:"racial"`
				Date   string `json:"date"`
				Id     string `json:"id"`
				Text   string `json:"text"`
			}
		}
	}
	Copyright string `json:"copyright"`
}

func Start() {

	// Creating new bot session
	goBot, err := discordgo.New("Bot " + config.Token)

	// Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Making our bot a user using User function .
	u, err := goBot.User("@me")

	// Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Storing our id from u to BotId .
	BotId = u.ID

	/* Adding handler function to handle our messages using AddHandler from
	discordgo package. We will declare messageHandler function later. */
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	// Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//If every thing works fine we will be printing this.
	fmt.Println("Bot is alive!")
}

//Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Bot musn't reply to it's own messages , to confirm it we perform this check.
	if m.Author.ID == BotId {
		return
	}
	// If we message ping to our bot in our discord it will return us pong .
	if m.Content == "joke" {
		res, err := http.Get(jokeAPI + "/jod")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var joke JokeApiResponse
		if err := json.Unmarshal(body, &joke); err != nil {
			fmt.Println("Couldn't unmarshal JSON")
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, joke.Contents.Jokes[0].Joke.Text)
	}

	// if strings.Contains(m.Content, "joke") {
	// 	fmt.Println("User needs a joke")
	// }
}
