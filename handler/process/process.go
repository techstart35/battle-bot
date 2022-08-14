package process

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
	"strings"
)

// 起動中のプロせセスを確認します
func ProcessHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content
	if cmd != shared.Command().Process {
		return
	}

	msg := make([]string, 0)

	for _, cID := range shared.GetProcess() {
		msg = append(msg, fmt.Sprintf("ChannelID: %s", cID))
	}

	if len(msg) == 0 {
		msg = append(msg, "実行中のプロセスはありません")
	}

	if err := shared.SendSimpleEmbedMessage(
		s, m.ChannelID, "実行中のプロセス", strings.Join(msg, "\n"),
	); err != nil {
		shared.SendErr(s, "実行中のプロセスメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}
}
