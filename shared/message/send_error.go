package message

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/guild"
	"log"
)

// エラーをTestチャンネルに送付します
//
// この中で発生したエラーはLogに出力します。
func SendErr(s *discordgo.Session, msg, guildID, channelID string, err error) {
	guildName, e := guild.GetGuildName(s, guildID)
	if e != nil {
		log.Println("ギルドIDを取得できません", e)
	}

	var sendErrTmpl = `
ギルド名: **%s**

チャンネル: %s

メッセージ: **%s**

継承したエラー: %s
`
	channelLink := shared.FormatChannelIDToLink(channelID)
	m := fmt.Sprintf(sendErrTmpl, guildName, channelLink, msg, err.Error())

	if e = SendSimpleEmbedMessage(s, AdminChannelID, "エラーが発生しました", m, shared.ColorRed); e != nil {
		errors.LogErr("エラーメッセージをAdminサーバーに送信できません", e)
	}
}
