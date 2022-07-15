package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/message"
	"github.com/techstart35/battle-bot/discord/shared"
	"log"
	"strings"
)

// 起動中のプロせセスを確認します
func ProcessHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content
	if cmd != ProcessCommand {
		return
	}

	msg := make([]string, 0)

	for k, v := range shared.IsProcessing {
		msg = append(msg, fmt.Sprintf("ChannelID: %s, Status: %v", k, v))
	}

	if len(msg) == 0 {
		msg = append(msg, "実行中のプロセスはありません")
	}

	if err := message.SendSimpleEmbedMessage(
		s, m.ChannelID, "実行中のプロセス", strings.Join(msg, "\n"),
	); err != nil {
		log.Println(err)
	}
}
