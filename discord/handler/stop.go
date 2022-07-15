package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/message"
	"github.com/techstart35/battle-bot/discord/shared"
	"log"
)

// 停止処理を実行します
func StopHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content
	if cmd != StopCommand {
		return
	}

	if _, ok := shared.IsProcessing[m.ChannelID]; !ok {
		if err := message.SendSimpleEmbedMessage(
			s, m.ChannelID, "キャンセル処理の実行", "このチャンネルで起動されたバトルはありません",
		); err != nil {
			log.Println(err)
		}

		return
	}

	// チャンネル一覧から削除
	delete(shared.IsProcessing, m.ChannelID)

	if err := message.SendSimpleEmbedMessage(
		s, m.ChannelID, "キャンセル処理の実行", "このチャンネルで起動されたバトルをキャンセルしました",
	); err != nil {
		log.Println(err)
	}
}
