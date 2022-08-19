package message

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/guild"
	"log"
)

// エラー送信のリクエストです
type SendErrReq struct {
	Message   string
	GuildID   string
	ChannelID string
	Err       error
}

// エラーをTestチャンネルに送付します
//
// この中で発生したエラーはLogに出力します。
func SendErr(s *discordgo.Session, req SendErrReq) {
	guildName, e := guild.GetGuildName(s, req.GuildID)
	if e != nil {
		log.Println("ギルドIDを取得できません", e)
	}

	var sendErrTmpl = `
ギルド名: **%s**

チャンネル: %s

メッセージ: **%s**

継承したエラー: %s
`
	channelLink := "none"
	if req.ChannelID != "none" {
		shared.FormatChannelIDToLink(req.ChannelID)
	}

	m := fmt.Sprintf(sendErrTmpl, guildName, channelLink, req.Message, req.Err.Error())

	embedInfo := &discordgo.MessageEmbed{
		Title:       "エラーが発生しました",
		Description: m,
		Color:       shared.ColorRed,
	}

	_, e = s.ChannelMessageSendEmbed(AdminChannelID, embedInfo)
	if e != nil {
		errors.LogErr("エラーメッセージをAdminサーバーに送信できません", e)
	}
}
