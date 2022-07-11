package noentry

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/shared"
)

var template = `
エントリーが無かったため、試合は開始されません
`

// NoEntryのメッセージを送信します
func SendNoEntryMessage(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	anotherChannelID string,
) error {
	// キャンセル指示を確認
	if !shared.IsProcessing[entryMessage.ChannelID] {
		return nil
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       "No Entry",
		Description: template,
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
