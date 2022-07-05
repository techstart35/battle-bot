package message

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Winnerのメッセージを送信します
func SendWinnerMessage(s *discordgo.Session, entryMessage *discordgo.Message, winner *discordgo.User) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       "👑 Winner 👑",
		Description: fmt.Sprintf("勝者：<@%s>", winner.ID),
		Color:       0xff0000,
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	msg, err := s.ChannelMessageSend(entryMessage.ChannelID, fmt.Sprintf("<@%s>さん、おめでとう🎉", winner.ID))
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	if err := s.MessageReactionAdd(msg.ChannelID, msg.ID, "🎉"); err != nil {
		return errors.New(fmt.Sprintf("リアクションを付与できません: %v", err))
	}

	return nil
}
