package entry

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/shared"
)

// 別チャンネルの指定がなかった場合のテンプレートです
var noAnotherChannelTemplate = `
⚔️ Battle Royale ⚔️

⚡️主催者：<@%s>
⚡️勝者：**1名**
⚡️エントリー：以下の⚔️にリアクション
⚡️開始：このメッセージ送信から**2分後**
`

// 別チャンネルの指定があった場合のテンプレートです
var withAnotherChannelTemplate = `
ねだるな！勝ち取れ🔥🔥

⚡️主催者：<@%s>
⚡️勝者：**1名**
⚡️エントリー：以下の⚔️にリアクション
⚡️開始：このメッセージ送信から**2分後**
⚡️配信チャンネル：<#%s>
`

// エントリーメッセージを送信します
//
// 起動元のチャンネルのみに送信します。
func SendEntryMessage(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
	anotherChannelID string,
) (*discordgo.Message, error) {
	if !shared.IsProcessing[m.ChannelID] {
		return nil, nil
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       "⚔️ Battle Royale ⚔️",
		Description: fmt.Sprintf(noAnotherChannelTemplate, m.Author.ID),
		Color:       0x0099ff,
	}

	// 別チャンネルの指定があった場合はテンプレートを差し替え
	if anotherChannelID != "" {
		embedInfo.Description = fmt.Sprintf(
			withAnotherChannelTemplate,
			m.Author.ID,
			anotherChannelID,
		)
	}

	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embedInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	// リアクションを付与
	if err := s.MessageReactionAdd(m.ChannelID, msg.ID, "⚔️"); err != nil {
		return nil, errors.New(fmt.Sprintf("リアクションを付与できません: %v", err))
	}

	return msg, nil
}
