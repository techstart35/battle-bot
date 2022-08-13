package noentry

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
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
			return shared.CreateErr("メッセージの送信に失敗しました", err)
		}
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return shared.CreateErr("メッセージの送信に失敗しました", err)
	}

	return nil
}
