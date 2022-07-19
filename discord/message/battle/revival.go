package battle

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/message/battle/template"
	"github.com/techstart35/battle-bot/discord/shared"
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
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	if anotherChannelID != "" {
		_, err := s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}
	}

	return nil
}
