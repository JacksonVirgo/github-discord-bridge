package discord

import (
	"fmt"

	"github.com/JacksonVirgo/github-discord-bridge/src/github/issues"
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
		messageOutsideThread(s, m, channel)
		return
	}

	messageInThread(s, m, channel)
}

func messageInThread(_ *discordgo.Session, m *discordgo.MessageCreate, channel *discordgo.Channel) {
	var channelId = channel.ParentID
	if channelId != discordContext.channelId {
		return
	}

	var threadTitle = channel.Name
	var threadNumber, threadNumErr = utils.ExtractIssueNumberFromThreadTitle(threadTitle)
	if threadNumErr != nil {
		return
	}

	var header = fmt.Sprintf("> Posted by **@%s**\n> <sub>%s</sub>\n\n", m.Author.Username, m.ID)
	var content = fmt.Sprintf("%s%s", header, m.Content)

	err := issues.CreateIssueComment(threadNumber, content)
	if err != nil {
		return
	}
}

func messageOutsideThread(s *discordgo.Session, m *discordgo.MessageCreate, _ *discordgo.Channel) {
	if m.Content == "!get-issues" {
		issues, err := issues.GetIssues()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}

		var returnStr = ""
		for _, issue := range issues {
			returnStr = returnStr + "> " + issue + "\n"
		}

		s.ChannelMessageSend(m.ChannelID, returnStr)
	} else if m.Content == "!sync-tags" {
		labels, err := issues.GetIssueLabels()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}

		forumChannel, err := s.State.Channel(discordContext.channelId)
		if err != nil {
			forumChannel, err = s.Channel(discordContext.channelId)
			if err != nil {
				return
			}
		}

		if forumChannel.Type != discordgo.ChannelTypeGuildForum {
			return
		}

		var existingTagsMap = make(map[string]bool)
		for _, tag := range forumChannel.AvailableTags {
			existingTagsMap[tag.Name] = true
		}

		var newTags []discordgo.ForumTag
		for _, label := range labels {
			if !existingTagsMap[label.Name] {
				newTags = append(newTags, discordgo.ForumTag{
					Name: label.Name,
				})
			}
		}

		if len(newTags) > 0 {
			updatedTags := append(forumChannel.AvailableTags, newTags...)
			_, err := s.ChannelEditComplex(discordContext.channelId, &discordgo.ChannelEdit{
				AvailableTags: &updatedTags,
			})
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Failed to update forum tags: "+err.Error())
				return
			}
			s.ChannelMessageSend(m.ChannelID, "Forum tags updated successfully.")
		} else {
			s.ChannelMessageSend(m.ChannelID, "All tags are already up to date.")
		}

		var tagsList = "Forum tags:\n"
		for _, tag := range forumChannel.AvailableTags {
			tagsList += "> " + tag.Name + "\n"
		}

		if tagsList == "Forum tags:\n" {
			tagsList = "No tags available in this forum channel."
		}

		s.ChannelMessageSend(m.ChannelID, tagsList)
	}
}
