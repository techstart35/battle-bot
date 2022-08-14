package stop

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
)

// 停止処理を実行します
func StopHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content
	if cmd != shared.Command().Stop {
		return
	}

	if !shared.IsProcessing(m.ChannelID) {
		if err := shared.SendSimpleEmbedMessage(
			s, m.ChannelID, "キャンセル処理の実行", "このチャンネルで起動されたバトルはありません",
		); err != nil {
			shared.SendErr(s, "起動されたバトルが無い場合のメッセージを送信できません", m.GuildID, m.ChannelID, err)
			return
		}

		return
	}

	// キャンセル処理を実行
	shared.CancelProcess(m.ChannelID)

	if err := shared.SendSimpleEmbedMessage(
		s, m.ChannelID, "キャンセル処理の実行", "このチャンネルで起動されたバトルをキャンセルしました",
	); err != nil {
		shared.SendErr(s, "キャンセル処理実行メッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}
}
