package message

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var noEntryTemplate = `
エントリーが無かったため、試合は開始されません
`

// NoEntryのメッセージを送信します
func SendNoEntryMessage(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	anotherChannelID string,
) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       "No Entry",
		Description: noEntryTemplate,
		Color:       0xff0000,
	}

	if anotherChannelID != "" {
		_, err := s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return nil
}
