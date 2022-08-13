package countdown

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
)

// エントリーチャンネルに送信するカウントダウンメッセージです
var entryChannelTemplate = `
開始まであと **%d秒**

⚔️-対戦
💥-自滅
☀️-敗者なし

Are You Ready?🔥🔥

<#%s> でも配信中 💬
`

// エントリーチャンネルに送信するカウントダウンメッセージです
//
// 別チャンネルが指定されていない場合に使用します。
var noAnotherChannelTemplate = `
開始まであと **%d秒**

⚔️-対戦
💥-自滅
☀️-敗者なし

Are You Ready?🔥🔥
`

// 別チャンネルに送信するカウントダウンメッセージです
var anotherChannelTemplate = `
開始まであと **%d秒**

⚔️-対戦
💥-自滅
☀️-敗者なし

Are You Ready?🔥🔥

▼エントリーはこちら
<#%s>
`

// カウントダウンメッセージを送信します
func SendCountDownMessage(
	s *discordgo.Session,
	entryMsg *discordgo.Message,
	beforeStart uint,
	anotherChannelID string,
) error {
	// キャンセル指示を確認
	if !shared.IsProcessing[entryMsg.ChannelID] {
		return nil
	}

	var color int
	switch beforeStart {
	case 60:
		color = 0x0099ff
	case 30:
		color = 0x3cb371
	case 10:
		color = 0xffd700
	}

	// 別チャンネルが無い場合を想定
	embedInfo := &discordgo.MessageEmbed{
		Title: "⚔️ Battle Royale ⚔️",
		Description: fmt.Sprintf(
			noAnotherChannelTemplate,
			beforeStart,
		),
		Color: color,
	}

	// 別チャンネルがあった場合
	if anotherChannelID != "" {
		// エントリーチャンネルに送信
		embedInfo.Description = fmt.Sprintf(
			entryChannelTemplate,
			beforeStart,
			anotherChannelID,
		)

		_, err := s.ChannelMessageSendEmbed(entryMsg.ChannelID, embedInfo)
		if err != nil {
			return shared.CreateErr("メッセージの送信に失敗しました", err)
		}

		// 別チャンネルに送信
		embedInfo.Description = fmt.Sprintf(
			anotherChannelTemplate,
			beforeStart,
			entryMsg.ChannelID,
		)

		_, err = s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return shared.CreateErr("メッセージの送信に失敗しました", err)
		}

		return nil
	}

	_, err := s.ChannelMessageSendEmbed(entryMsg.ChannelID, embedInfo)
	if err != nil {
		return shared.CreateErr("メッセージの送信に失敗しました", err)
	}

	return nil
}
