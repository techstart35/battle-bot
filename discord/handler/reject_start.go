package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/message"
	"github.com/techstart35/battle-bot/discord/shared"
	"log"
)

// 新規起動を禁止します
//
// sharedのIsStartRejectedをtrueに変更します。
func RejectStartHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content
	if cmd != RejectStartCommand {
		return
	}

	shared.IsStartRejected = true

	if err := message.SendSimpleEmbedMessage(
		s, m.ChannelID, "新規起動の停止", "新規起動を停止しました。",
	); err != nil {
		log.Println(err)
	}
}
