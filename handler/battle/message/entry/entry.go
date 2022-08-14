package entry

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
)

// 別チャンネルの指定がなかった場合のテンプレートです
var noAnotherChannelTemplate = `
⚡️主催者：<@%s>
⚡️勝者：**1名**
⚡️エントリー：以下の⚔️にリアクション
⚡️開始：このメッセージ送信から**2分後**
`

// 別チャンネルの指定があった場合のテンプレートです
var withAnotherChannelTemplate = `
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
	if shared.IsCanceled(m.GuildID) {
		return nil, nil
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       "⚔️ Battle Royale ⚔️",
		Description: fmt.Sprintf(noAnotherChannelTemplate, m.Author.ID),
		Color:       shared.ColorBlue,
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
		return nil, shared.CreateErr("メッセージの送信に失敗しました", err)
	}

	// リアクションを付与
	if err := s.MessageReactionAdd(m.ChannelID, msg.ID, "⚔️"); err != nil {
		return nil, shared.CreateErr("リアクションを付与できません", err)
	}

	return msg, nil
}
