package reject_start

//
//import (
//	"github.com/bwmarrin/discordgo"
//	"github.com/techstart35/battle-bot/shared"
//	"github.com/techstart35/battle-bot/shared/message"
//)
//
//// 新規起動を禁止します
////
//// sharedのIsStartRejectedをtrueに変更します。
//func RejectStartHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
//	cmd := m.Content
//	if cmd != shared.Command().RejectStart {
//		return
//	}
//
//	shared.IsStartRejected = true
//
//	if err := message.SendSimpleEmbedMessage(
//		s, m.ChannelID, "新規起動の停止", "新規起動を停止しました。", 0,
//	); err != nil {
//		message.SendErr(s, "新規起動の停止メッセージを送信できません", m.GuildID, m.ChannelID, err)
//		return
//	}
//}
