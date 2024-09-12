package controllers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func StartDiscordBot(token string) (*discordgo.Session, error) {
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return nil, err
	}

	bot.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds | discordgo.IntentsMessageContent
	bot.AddHandler(onReady)
	bot.AddHandler(messageCreate)
	bot.AddHandler(threadCreate)

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

func onReady(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Printf("Logged in as: %v\n", s.State.User.Username)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Content == "!get-issues" {
		issues, err := GetIssues()
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

func threadCreate(s *discordgo.Session, t *discordgo.ThreadCreate) {
	if t.NewlyCreated {
		message_id := t.LastMessageID
		message, err := s.ChannelMessage(t.ID, message_id)
		if err != nil {
			fmt.Println(err)
			return
		}

		var header = fmt.Sprintf("> Posted by **@%s**\n\n", message.Author.Username)
		var content = fmt.Sprintf("%s%s", header, message.Content)

		var issue, create_err = CreateIssue(CreateIssueRequest{
			Title:  t.Name,
			Body:   content,
			Labels: []string{},
			Headers: Headers{
				XGitHubApiVersion: "2022-11-28",
			},
			Owner: githubContext.author,
			Repo:  githubContext.repo,
		})

		if create_err != nil {
			s.ChannelMessageSend(t.ID, create_err.Error())
			return
		}

		var threadRename = fmt.Sprintf("%d) %s", issue.Number, t.Name)
		_, err = s.ChannelEdit(t.ID, &discordgo.ChannelEdit{
			Name: threadRename,
		})

		if err != nil {
			s.ChannelMessageSend(t.ID, fmt.Sprintf("Failed to rename thread: %s", err.Error()))
			return
		}

		var response = fmt.Sprintf("[Issue #%d created](<%s>)", issue.Number, issue.HtmlUrl)
		s.ChannelMessageSend(t.ID, response)
	}
}
