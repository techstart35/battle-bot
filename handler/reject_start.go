package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
)

// 新規起動を禁止します
//
// sharedのIsStartRejectedをtrueに変更します。
func RejectStartHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content
	if cmd != shared.Command().RejectStart {
		return
	}

	shared.IsStartRejected = true

	if err := shared.SendSimpleEmbedMessage(
		s, m.ChannelID, "新規起動の停止", "新規起動を停止しました。",
	); err != nil {
		shared.SendErr(s, "新規起動の停止メッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}
}
