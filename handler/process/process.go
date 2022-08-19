package process

//
//import (
//	"fmt"
//	"github.com/bwmarrin/discordgo"
//	"github.com/techstart35/battle-bot/shared"
//	"github.com/techstart35/battle-bot/shared/guild"
//	"github.com/techstart35/battle-bot/shared/message"
//	"strings"
//)
//
//// 起動中のプロせセスを確認します
//func ProcessHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
//	cmd := m.Content
//	if cmd != shared.Command().Process {
//		return
//	}
//
//	msg := make([]string, 0)
//
//	for guildID, isProcessing := range shared.GetStore() {
//		status := "キャンセル済"
//		if isProcessing {
//			status = "起動中"
//		}
//
//		guildName, err := guild.GetGuildName(s, guildID)
//		if err != nil {
//			message.SendErr(s, "ギルドIDを取得できません", m.GuildID, m.ChannelID, err)
//			return
//		}
//
//		msg = append(msg, fmt.Sprintf("%s｜サーバー名: %s", status, guildName))
//	}
//
//	if len(msg) == 0 {
//		msg = append(msg, "実行中のストアはありません")
//	}
//
//	if err := message.SendSimpleEmbedMessage(
//		s, m.ChannelID, "実行中のストア", strings.Join(msg, "\n"), 0,
//	); err != nil {
//		message.SendErr(s, "実行中のストアメッセージを送信できません", m.GuildID, m.ChannelID, err)
//		return
//	}
//}
