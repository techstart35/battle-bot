package message

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var entryTemplate = `
ねだるな！勝ち取れ🔥🔥

⚡️主催者：<@%s>
⚡️勝者：**1名**
⚡️エントリー：以下の⚔️にリアクション
⚡️開始：このメッセージ送信から**2分後**
`

var entryTemplateWithAnotherChannel = `
ねだるな！勝ち取れ🔥🔥

⚡️主催者：<@%s>
⚡️勝者：**1名**
⚡️エントリー：以下の⚔️にリアクション
⚡️開始：このメッセージ送信から**2分後**
⚡️配信チャンネル：<#%s>
`

// エントリーメッセージを送信します
//
// 引数のチャンネルIDがある場合、そちらにもメッセージを送信します。
func SendEntryMessage(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
	anotherChannelID string,
) (*discordgo.Message, error) {
	embedInfo := &discordgo.MessageEmbed{
		Title:       "⚔️ Giveaway Battle ⚔️",
		Description: fmt.Sprintf(entryTemplate, m.Author.ID),
		Color:       0x0099ff,
	}

	if anotherChannelID != "" {
		embedInfo.Description = fmt.Sprintf(
			entryTemplateWithAnotherChannel,
			m.Author.ID,
			anotherChannelID,
		)
	}

	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embedInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	if err := s.MessageReactionAdd(m.ChannelID, msg.ID, "⚔️"); err != nil {
		return nil, errors.New(fmt.Sprintf("リアクションを付与できません: %v", err))
	}

	return msg, nil
}
