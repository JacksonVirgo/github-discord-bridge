package discord

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type DiscordContext struct {
	token     string
	channelId string
}

var discordContext = &DiscordContext{}

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

func StartDiscordBot() (*discordgo.Session, error) {
	token := discordContext.token
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return nil, err
	}

	bot.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds | discordgo.IntentsMessageContent
	bot.AddHandler(OnReady)
	bot.AddHandler(MessageCreate)
	bot.AddHandler(ThreadCreate)
	bot.AddHandler(ThreadUpdate)

	err = bot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return nil, err
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Close()
	return bot, nil
}
