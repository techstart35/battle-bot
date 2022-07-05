package message

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// シンプルな埋め込みメッセージを送信します
func SendSimpleEmbedMessage(s *discordgo.Session, channelID, title, description string) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       0x000000,
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return nil
}
