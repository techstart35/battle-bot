package noentry

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
)

var template = `
エントリーが無かったため、試合は開始されません
`

// NoEntryのメッセージを送信します
func SendNoEntryMessage(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
	anotherChannelID string,
) error {
	// キャンセル指示を確認
	if shared.IsCanceled(m.GuildID) {
		return nil
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       "No Entry",
		Description: template,
		Color:       shared.ColorRed,
	}

	if anotherChannelID != "" {
		_, err := s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embedInfo)
	if err != nil {
		return errors.NewError("メッセージの送信に失敗しました", err)
	}

	return nil
}
