package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Printf("Logged in as: %v\n", s.State.User.Username)
}
