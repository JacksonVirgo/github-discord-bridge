package discord

import (
	"fmt"

	"github.com/JacksonVirgo/github-discord-bridge/src/github"
	"github.com/JacksonVirgo/github-discord-bridge/src/github/issues"
	"github.com/bwmarrin/discordgo"
)

func ThreadCreate(s *discordgo.Session, t *discordgo.ThreadCreate) {
	if t.ParentID != discordContext.channelId {
		return
	}

	if t.NewlyCreated {
		messages, err := getAllMessagesFromThread(s, t.ID)
		if err != nil {
			fmt.Printf("Message Err: %s", err)
			return
		}
		if len(messages) == 0 {
			fmt.Printf("Message Err: No Messages")
			return
		}

		oldestMessage := messages[0]
		for _, message := range messages {
			if message.Timestamp.Before(oldestMessage.Timestamp) {
				oldestMessage = message
			}
		}

		var header = fmt.Sprintf("> Posted by **@%s**\n\n", oldestMessage.Author.Username)
		var content = fmt.Sprintf("%s%s", header, oldestMessage.Content)

		var issue, create_err = issues.CreateIssue(issues.CreateIssueRequest{
			Title:  t.Name,
			Body:   content,
			Labels: []string{},
			Headers: issues.Headers{
				XGitHubApiVersion: "2022-11-28",
			},
			Owner: github.GetAuthor(),
			Repo:  github.GetRepo(),
		})

		if create_err != nil {
			fmt.Printf("Issue Creation Err: %s", create_err.Error())
			s.ChannelMessageSend(t.ID, create_err.Error())
			return
		}

		var threadRename = fmt.Sprintf("%d) %s", issue.Number, t.Name)
		_, err = s.ChannelEdit(t.ID, &discordgo.ChannelEdit{
			Name: threadRename,
		})

		if err != nil {
			fmt.Printf("Thread Rename Err: %s", err.Error())
			s.ChannelMessageSend(t.ID, fmt.Sprintf("Failed to rename thread: %s", err.Error()))
			return
		}

		var response = fmt.Sprintf("[Issue #%d created](<%s>)", issue.Number, issue.HtmlUrl)
		s.ChannelMessageSend(t.ID, response)
	}
}

func getAllMessagesFromThread(s *discordgo.Session, threadID string) ([]*discordgo.Message, error) {
	var allMessages []*discordgo.Message
	var lastMessageID string

	for {
		messages, err := s.ChannelMessages(threadID, 5, lastMessageID, "", "")
		if err != nil {
			return nil, err
		}

		if len(messages) == 0 {
			break
		}
		allMessages = append(allMessages, messages...)
		lastMessageID = messages[len(messages)-1].ID
	}

	return allMessages, nil
}
