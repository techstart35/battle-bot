package discord

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// バトルメッセージを送信します
func sendBattleMessage(
	s *discordgo.Session,
	users []*discordgo.User,
	entryMessage *discordgo.Message,
	channelID string,
) error {
	if len(users) == 0 {
		if err := sendNoEntryMessage(s, entryMessage); err != nil {
			return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}

		return nil
	}

	var survivor []*discordgo.User
	survivor = users

	var battleMessages []string

	for {
		if len(survivor) == 1 {
			break
		}

		// 2人抽出して、テンプレートメッセージにはめる
		usr1 := survivor[0]

		// メッセージをスライスに格納
		msg := fmt.Sprintf(sampleTmpl, usr1.Username)
		battleMessages = append(battleMessages, msg)
	}

	// 格納したメッセージを改行して送信
	embedInfo := &discordgo.MessageEmbed{
		Title:       "⚔️ Battle ⚔️",
		Description: strings.Join(battleMessages, ""),
		Color:       0xff0000,
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return nil
}

var sampleTmpl = "%s は倒された"

// NoEntryのメッセージを送信します
func sendNoEntryMessage(s *discordgo.Session, entryMessage *discordgo.Message) error {
	// 格納したメッセージを改行して送信
	embedInfo := &discordgo.MessageEmbed{
		Title:       "No Entry",
		Description: "エントリーが無かったため、試合は開始されません",
		Color:       0xff0000,
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return nil
}
