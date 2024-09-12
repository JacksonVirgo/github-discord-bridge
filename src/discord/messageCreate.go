package discord

import (
	"fmt"

	"github.com/JacksonVirgo/github-discord-bridge/src/github"
	"github.com/JacksonVirgo/github-discord-bridge/src/utils"
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		channel, err = s.Channel(m.ChannelID)
		if err != nil {
			return
		}
	}

	if !utils.CheckChannelIsThread(channel) {
		return
	}

	var channelId = channel.ParentID
	if channelId != discordContext.channelId {
		return
	}

	var threadTitle = channel.Name
	var threadNumber, threadNumErr = utils.ExtractIssueNumberFromThreadTitle(threadTitle)
	if threadNumErr != nil {
		return
	}

	var header = fmt.Sprintf("> Posted by **@%s**\n\n", m.Author.Username)
	var content = fmt.Sprintf("%s%s", header, m.Content)

	err = github.CreateIssueComment(threadNumber, content)
	if err != nil {
		return
	}

	if m.Content == "!get-issues" {
		issues, err := github.GetIssues()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}

		var returnStr = ""
		for _, issue := range issues {
			returnStr = returnStr + "> " + issue + "\n"
		}

		s.ChannelMessageSend(m.ChannelID, returnStr)
	}
}
