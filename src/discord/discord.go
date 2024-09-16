package discord

import (
	"errors"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

type DiscordContext struct {
	token     string
	channelId string
}

var discordContext = &DiscordContext{}

var bot *discordgo.Session

func LoadDiscordContext() error {
	discordToken := os.Getenv("DISCORD_TOKEN")
	channelId := os.Getenv("DISCORD_CHANNEL_ID")

	if discordToken == "" || channelId == "" {
		return errors.New("missing environment variables")
	}

	*discordContext = DiscordContext{
		token:     discordToken,
		channelId: channelId,
	}

	return nil
}

func StartDiscordBot() error {
	token := discordContext.token
	var err error
	bot, err = discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return err
	}

	bot.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds | discordgo.IntentsMessageContent
	bot.AddHandler(OnReady)
	bot.AddHandler(MessageCreate)
	bot.AddHandler(ThreadCreate)
	bot.AddHandler(ThreadUpdate)

	err = bot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return err
	}

	return nil
}

func CloseDiscordBot() {
	if bot != nil {
		bot.Close()
	}
}
