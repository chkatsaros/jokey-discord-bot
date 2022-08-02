package bot

import (
	"fmt"
	"jokey-discord-bot/config"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/icelain/jokeapi"
)

var BotId string

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
	fmt.Println("Jokey is up and running!")
}

/* Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s,
second one is discordgo.MessageCreate which is m. */
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Bot musn't reply to it's own messages, to confirm it we perform this check.
	if m.Author.ID == BotId {
		return
	}

	if strings.ToLower(m.Content) == "jokey help" {
		helpMssg := "How can I help you today?"
		_, _ = s.ChannelMessageSend(m.ChannelID, helpMssg)
	}

	if strings.ToLower(m.Content) == "joke" {

		joke := getRandomJoke()

		var complexMssg discordgo.MessageSend

		complexMssg.Components = []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.SelectMenu{
						// Select menu, as other components, must have a customID, so we set it to this value.
						CustomID:    "select",
						Placeholder: "Choose your favorite programming language 👇",
						Options: []discordgo.SelectMenuOption{
							{
								Label: "Go",
								// As with components, this things must have their own unique "id" to identify which is which.
								// In this case such id is Value field.
								Value: "go",
								Emoji: discordgo.ComponentEmoji{
									Name: "🦦",
								},
								// You can also make it a default option, but in this case we won't.
								Default:     false,
								Description: "Go programming language",
							},
							{
								Label: "JS",
								Value: "js",
								Emoji: discordgo.ComponentEmoji{
									Name: "🟨",
								},
								Description: "JavaScript programming language",
							},
							{
								Label: "Python",
								Value: "py",
								Emoji: discordgo.ComponentEmoji{
									Name: "🐍",
								},
								Description: "Python programming language",
							},
						},
					},
				},
			},
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, joke)
		_, _ = s.ChannelMessageSendComplex(m.ChannelID, &complexMssg)
	}
}

func getRandomJoke() string {
	ctgs := []string{"Programming", "Dark", "Miscellaneous", "Pun", "Spooky"}

	api := jokeapi.New()

	api.Set(jokeapi.Params{Categories: ctgs})
	res, err := api.Fetch()

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	switch res.JokeType {
	case "single":
		return res.Joke[0]
	case "twopart":
		return res.Joke[0] + "\n" + res.Joke[1]
	}

	return ""
}
