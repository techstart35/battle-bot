package message

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
)

// シンプルな埋め込みメッセージのリクエストです
type SendSimpleEmbedMessageReq struct {
	ChannelID string
	Title     string
	Content   string
	ColorCode int
}

// シンプルな埋め込みメッセージを送信します
func SendSimpleEmbedMessage(s *discordgo.Session, req SendSimpleEmbedMessageReq) error {
	col := shared.ColorBlack
	if req.ColorCode != 0 {
		col = req.ColorCode
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       req.Title,
		Description: req.Content,
		Color:       col,
	}

	_, err := s.ChannelMessageSendEmbed(req.ChannelID, embedInfo)
	if err != nil {
		return errors.NewError("メッセージの送信に失敗しました", err)
	}

	return nil
}
