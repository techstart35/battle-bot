package message

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
)

// シンプルな埋め込みメッセージを送信します
func SendSimpleEmbedMessage(s *discordgo.Session, channelID, title, description string, color int) error {
	col := shared.ColorBlack
	if color != 0 {
		col = color
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       col,
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embedInfo)
	if err != nil {
		return errors.NewError("メッセージの送信に失敗しました", err)
	}

	return nil
}
