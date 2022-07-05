package message

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/shared"
	"strings"
)

var startTemplate = `
⚡️挑戦者（%d名）：%s
⚡️勝者：1名
`

var startTemplateWithAnotherChannel = `
⚡️挑戦者(%d名）：%s
⚡️勝者：1名
⚡️<#%s> チャンネルでも配信中💬
`

// 開始メッセージを送信します
func SendStartMessage(
	s *discordgo.Session,
	entryMsg *discordgo.Message,
	anotherChannelID string,
) ([]*discordgo.User, error) {
	users, err := shared.GetReactedUsers(s, entryMsg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("リアクションしたユーザーの取得に失敗しました: %v", err))
	}

	var challengers []string
	for _, v := range users {
		challengers = append(challengers, fmt.Sprintf("<@%s>", v.ID))
	}

	userStr := strings.Join(challengers, " ")

	embedInfo := &discordgo.MessageEmbed{
		Title:       "⚔️ Battle Start ⚔️",
		Description: fmt.Sprintf(startTemplate, len(challengers), userStr),
		Color:       0xff0000,
	}

	// チャンネルIDが入っている場合は、別チャンネルに送信 & Descriptionの書き換えを行います。
	if anotherChannelID != "" {
		_, err := s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}

		embedInfo.Description = fmt.Sprintf(
			startTemplateWithAnotherChannel,
			len(challengers),
			userStr,
			anotherChannelID,
		)
	}

	_, err = s.ChannelMessageSendEmbed(entryMsg.ChannelID, embedInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return users, nil
}
