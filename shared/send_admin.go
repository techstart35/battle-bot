package shared

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

const AdminChannelID = "1003130506881277952"

// 開始時に自分のサーバーにメッセージを送信します
func SendStartMessageToAdmin(s *discordgo.Session, guildID, channelID string, command []string) error {
	guildName := guildID
	if name, ok := GuildName[guildID]; ok {
		guildName = name
	}

	var template = `

**サーバー名**
%s

**起動チャンネル**
%s

**実行コマンド**
%s

**開始時刻**
%s
`

	channelLink := FormatChannelIDToLink(channelID)
	now := time.Now().Format("2006-01-02 15:04:05")

	msg := fmt.Sprintf(template, guildName, channelLink, strings.Join(command, " "), now)
	if err := SendSimpleEmbedMessage(s, AdminChannelID, "起動通知", msg, ColorCyan); err != nil {
		return CreateErr("起動通知メッセージを送信できません", err)
	}

	return nil
}

// 停止コマンド実行時に自分のサーバーにメッセージを送信します
func SendStopMessageToAdmin(s *discordgo.Session, guildID string) error {
	guildName := guildID

	if name, ok := GuildName[guildID]; ok {
		guildName = name
	}

	var template = `
**サーバー名**
%s

**停止時間**
%s
`

	now := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(template, guildName, now)
	if err := SendSimpleEmbedMessage(s, AdminChannelID, "停止コマンド通知", msg, ColorYellow); err != nil {
		return CreateErr("停止コマンド実行通知メッセージを送信できません", err)
	}

	return nil
}
