package battle

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/shared"
)

// 復活のテンプレートをランダムに取得します
func GetRandomRevivalTmpl(user *discordgo.User) string {
	var tmpl = []string{
		fmt.Sprintf("⚰️｜** %s ** は穢土転生により復活した。", user.Username),
	}

	return tmpl[shared.RandInt(1, len(tmpl)+1)-1]
}

// 復活メッセージを送信します
func SendRevivalMessage(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	user *discordgo.User,
	anotherChannelID string,
) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       "敗者復活🔥",
		Description: GetRandomRevivalTmpl(user),
		Color:       0xffc0cb,
	}

	if anotherChannelID != "" {
		_, err := s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return nil
}
