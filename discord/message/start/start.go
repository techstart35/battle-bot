package start

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/shared"
	"math"
	"strings"
)

// エントリーチャンネルに送信するメッセージです
var entryChannelTemplate = `
⚡️挑戦者(%d名）：%s
⚡️勝者：**1名**
⚡️勝率：**%v％**
⚡️<#%s> チャンネルでも配信中 💬
`

// エントリーチャンネルに送信するメッセージです
var entryChannelNoAnotherChannelTemplate = `
⚡️挑戦者(%d名）：%s
⚡️勝者：**1名**
⚡️勝率：**%v％**
`

// 別チャンネルに送信するメッセージです
//
// 別チャンネルを指定していない場合のエントリーチャンネルもこちらのテンプレートを使用します。
var anotherChannelTemplate = `
⚡️挑戦者（%d名）：%s
⚡️勝者：**1名**
⚡️勝率：**%v％**
`

// 開始メッセージを送信します
func SendStartMessage(
	s *discordgo.Session,
	entryMsg *discordgo.Message,
	anotherChannelID string,
) ([]*discordgo.User, error) {
	// キャンセル指示を確認
	if !shared.IsProcessing[entryMsg.ChannelID] {
		return nil, nil
	}

	users, err := shared.GetReactedUsers(s, entryMsg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("リアクションしたユーザーの取得に失敗しました: %v", err))
	}

	var challengers []string
	for _, v := range users {
		challengers = append(challengers, v.Username)
	}

	userStr := "100名を超えたため省略"
	if len(challengers) < 100 {
		userStr = strings.Join(challengers, ", ")
	}

	p := 1 / float64(len(challengers)) * 100
	probability := math.Round(p*10) / 10

	// 別チャンネルがない場合を想定
	embedInfo := &discordgo.MessageEmbed{
		Title: "⚔️ Battle Start ⚔️",
		Description: fmt.Sprintf(
			entryChannelNoAnotherChannelTemplate,
			len(challengers),
			userStr,
			probability,
		),
		Color: 0xff0000,
	}

	// 別チャンネルがあった場合
	if anotherChannelID != "" {
		// エントリーチャンネルに送信
		embedInfo.Description = fmt.Sprintf(
			entryChannelTemplate,
			len(challengers),
			userStr,
			probability,
			anotherChannelID,
		)

		_, err = s.ChannelMessageSendEmbed(entryMsg.ChannelID, embedInfo)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}

		// 別チャンネルに送信
		embedInfo.Description = fmt.Sprintf(
			anotherChannelTemplate,
			len(challengers),
			userStr,
			probability,
		)

		_, err = s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}

		return users, nil
	}

	_, err = s.ChannelMessageSendEmbed(entryMsg.ChannelID, embedInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return users, nil
}
