package battle

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/handler/battle/message/battle/template"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
)

// 復活メッセージを送信します
func SendRevivalMessage(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	user *discordgo.User,
	anotherChannelID string,
) error {
	// キャンセル指示を確認
	if shared.IsCanceled(entryMessage.GuildID) {
		return nil
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       "敗者復活",
		Description: template.GetRandomRevivalTmpl(user),
		Color:       shared.ColorPink,
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return errors.NewError("メッセージの送信に失敗しました", err)
	}

	if anotherChannelID != "" {
		_, err := s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}
	}

	return nil
}
