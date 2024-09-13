package discord

import (
	"fmt"
	"strings"

	"github.com/JacksonVirgo/github-discord-bridge/src/github/issues"
	"github.com/JacksonVirgo/github-discord-bridge/src/utils"
	"github.com/bwmarrin/discordgo"
)

func ThreadUpdate(s *discordgo.Session, t *discordgo.ThreadUpdate) {
	fmt.Printf("Thread Update: %v - %v\n", t.ParentID, discordContext.channelId)
	if t.ParentID != discordContext.channelId {
		return
	}

	var threadTitle = t.Name
	var threadNumber, threadNumErr = utils.ExtractIssueNumberFromThreadTitle(threadTitle)
	if threadNumErr != nil {
		fmt.Printf("Thread Update Err: %s", threadNumErr)
		return
	}

	// before := t.BeforeUpdate
	forumChannel, err := s.State.Channel(t.ParentID)
	if err != nil {
		forumChannel, err = s.Channel(t.ParentID)
		if err != nil {
			fmt.Printf("Forum Channel Err: %s", err)
			return
		}
	}

	oldTagIds := t.BeforeUpdate.AppliedTags
	newTagIds := t.AppliedTags
	if utils.UnorderedEqual(oldTagIds, newTagIds) {
		return
	}

	var newTagNames []string
	for _, appliedTagID := range t.AppliedTags {
		for _, availableTag := range forumChannel.AvailableTags {
			if appliedTagID == availableTag.ID {
				newTagNames = append(newTagNames, availableTag.Name)
				break
			}
		}
	}

	var oldTagNames []string
	for _, appliedTagID := range t.BeforeUpdate.AppliedTags {
		for _, availableTag := range forumChannel.AvailableTags {
			if appliedTagID == availableTag.ID {
				oldTagNames = append(oldTagNames, availableTag.Name)
				break
			}
		}
	}

	addedtags := utils.Difference(newTagNames, oldTagNames)
	removedTags := utils.Difference(oldTagNames, newTagNames)

	var response string = "```diff\n[ Updated Tags ]\n"
	if len(addedtags) > 0 {
		response += fmt.Sprintf("+ %v\n", strings.Join(addedtags, ", "))
	}
	if len(removedTags) > 0 {
		response += fmt.Sprintf("- %v\n", strings.Join(removedTags, ", "))
	}

	response += "```"

	err = issues.SetIssueLabels(threadNumber, newTagNames)
	if err != nil {
		fmt.Printf("Thread Update Err: %s", err)
		return
	}

	s.ChannelMessageSend(t.ID, response)
}
