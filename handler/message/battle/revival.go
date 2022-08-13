package battle

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/handler/message/battle/template"
	"github.com/techstart35/battle-bot/shared"
)

// 復活メッセージを送信します
func SendRevivalMessage(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	user *discordgo.User,
	anotherChannelID string,
) error {
	// キャンセル指示を確認
	if !shared.IsProcessing[entryMessage.ChannelID] {
		return nil
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       "敗者復活",
		Description: template.GetRandomRevivalTmpl(user),
		Color:       0xff69b4,
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return shared.CreateErr("メッセージの送信に失敗しました", err)
	}

	if anotherChannelID != "" {
		_, err := s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return shared.CreateErr("メッセージの送信に失敗しました", err)
		}
	}

	return nil
}
