package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ExtractIssueNumberFromThreadTitle(title string) (int, error) {
	sepIndex := strings.Index(title, ")")
	if sepIndex == -1 {
		return 0, fmt.Errorf("invalid format: missing closing parenthesis")
	}
	beforeParenthesis := title[:sepIndex]
	beforeParenthesis = strings.TrimSpace(beforeParenthesis)
	x, err := strconv.Atoi(beforeParenthesis)
	if err != nil {
		return 0, fmt.Errorf("invalid format: %v", err)
	}

	return x, nil
}

func CheckChannelIsThread(channel *discordgo.Channel) bool {
	return channel.Type == discordgo.ChannelTypeGuildPublicThread || channel.Type == discordgo.ChannelTypeGuildPrivateThread || channel.Type == discordgo.ChannelTypeGuildNewsThread
}
