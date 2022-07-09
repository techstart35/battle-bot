package message

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var countDownTemplate = `
開始まであと **%d秒**

Are You Ready?

▼エントリーはこちら
<#%s>
`

var entryChannelCountDownTemplate = `
開始まであと **%d秒**

Are You Ready?

<#%s> でも配信中 💬
`

// カウントダウンメッセージを送信します
func SendCountDownMessage(
	s *discordgo.Session,
	entryMsg *discordgo.Message,
	beforeStart uint,
	anotherChannelID string,
) error {
	var color int
	switch beforeStart {
	case 60:
		color = 0x0099ff
	case 30:
		color = 0x3cb371
	case 10:
		color = 0xffd700
	}

	embedInfo := &discordgo.MessageEmbed{
		Title: "⚔️ Giveaway Battle ⚔️",
		Description: fmt.Sprintf(
			countDownTemplate,
			beforeStart,
			entryMsg.ChannelID,
		),
		Color: color,
	}

	// チャンネルIDが入っている場合は、別チャンネルに送信 & Descriptionの書き換えを行います。
	if anotherChannelID != "" {
		_, err := s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}

		embedInfo.Description = fmt.Sprintf(
			entryChannelCountDownTemplate,
			beforeStart,
			anotherChannelID,
		)
	}

	_, err := s.ChannelMessageSendEmbed(entryMsg.ChannelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return nil
}
