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

	if !shared.IsProcessing(m.GuildID) {
		if err := shared.SendSimpleEmbedMessage(
			s, m.ChannelID, "キャンセル処理の実行", "このサーバーで起動されたバトルが無いか、キャンセル済みとなっています。",
		); err != nil {
			shared.SendErr(s, "起動されたバトルが無い場合のメッセージを送信できません", m.GuildID, m.ChannelID, err)
			return
		}

		return
	}

	// Adminサーバーに停止処理実行メッセージを送信します
	//
	// Notice: ここでエラーが発生しても処理は継続させます
	if err := shared.SendStopMessageToAdmin(s, m.GuildID); err != nil {
		shared.SendErr(s, "停止通知をAdminサーバーに送信できません", m.GuildID, m.ChannelID, err)
	}

	// キャンセル処理を実行
	shared.CancelProcess(m.GuildID)

	msg := `
このサーバーで起動されたバトルをキャンセルしました。
（反映まで最大1分かかります）
`

	if err := shared.SendSimpleEmbedMessage(
		s, m.ChannelID, "キャンセル処理の実行", msg,
	); err != nil {
		shared.SendErr(s, "キャンセル処理実行メッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}
}
