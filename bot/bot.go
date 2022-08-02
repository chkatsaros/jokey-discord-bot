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

		complexMssg.Content = joke
		complexMssg.Components = []discordgo.MessageComponent{
			// ActionRow is a container of all buttons within the same row.
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						// Label is what the user will see on the button.
						Label: "Yes",
						// Style provides coloring of the button. There are not so many styles tho.
						Style: discordgo.SuccessButton,
						// Disabled allows bot to disable some buttons for users.
						Disabled: false,
						// CustomID is a thing telling Discord which data to send when this button will be pressed.
						CustomID: "fd_yes",
					},
					discordgo.Button{
						Label:    "No",
						Style:    discordgo.DangerButton,
						Disabled: false,
						CustomID: "fd_no",
					},
					discordgo.Button{
						Label:    "I don't know",
						Style:    discordgo.LinkButton,
						Disabled: false,
						// Link buttons don't require CustomID and do not trigger the gateway/HTTP event
						URL: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
						Emoji: discordgo.ComponentEmoji{
							Name: "🤷",
						},
					},
				},
			},
		}

		// _, _ = s.ChannelMessageSend(m.ChannelID, joke)
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
