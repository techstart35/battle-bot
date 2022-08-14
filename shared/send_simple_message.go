package shared

import (
	"github.com/bwmarrin/discordgo"
)

// シンプルな埋め込みメッセージを送信します
func SendSimpleEmbedMessage(s *discordgo.Session, channelID, title, description string, color int) error {
	col := ColorBlack
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
		return CreateErr("メッセージの送信に失敗しました", err)
	}

	return nil
}
