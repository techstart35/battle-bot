package winner

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
)

// エントリーチャンネルに送信するメッセージです
var entryChannelTemplate = `
勝者：<@%s>
`

// 別チャンネルに送信するメッセージです
var anotherChannelTemplate = `
勝者：<@%s>

おめでとうございます🎉
`

// Winnerのメッセージを送信します
func SendWinnerMessage(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	winner *discordgo.User,
	anotherChannelID string,
) error {
	// キャンセル指示を確認
	if !shared.IsProcessing[entryMessage.ChannelID] {
		return nil
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       "👑 Winner 👑",
		Description: fmt.Sprintf(entryChannelTemplate, winner.ID),
		Color:       0xff0000,
	}

	// エントリーチャンネルにメッセージを送信
	{
		_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
		if err != nil {
			return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}

		msg, err := s.ChannelMessageSend(
			entryMessage.ChannelID,
			fmt.Sprintf("<@%s>さん、おめでとうございます🎉", winner.ID),
		)
		if err != nil {
			return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}

		if err := s.MessageReactionAdd(msg.ChannelID, msg.ID, "🎉"); err != nil {
			return errors.New(fmt.Sprintf("リアクションを付与できません: %v", err))
		}
	}

	// 別チャンネルにメッセージを送信
	if anotherChannelID != "" {
		ei := &discordgo.MessageEmbed{
			Title:       "👑 Winner 👑",
			Description: fmt.Sprintf(anotherChannelTemplate, winner.ID),
			Color:       0xff0000,
		}

		msg, err := s.ChannelMessageSendEmbed(anotherChannelID, ei)
		if err != nil {
			return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}

		if err := s.MessageReactionAdd(msg.ChannelID, msg.ID, "🎉"); err != nil {
			return errors.New(fmt.Sprintf("リアクションを付与できません: %v", err))
		}
	}

	return nil
}
