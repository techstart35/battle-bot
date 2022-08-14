package shared

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

const AdminChannelID = "1003130506881277952"

// 開始時に自分のサーバーにメッセージを送信します
func SendStartMessageToAdmin(s *discordgo.Session, guildID string, command []string) error {
	guildName := guildID
	if name, ok := GuildName[guildID]; ok {
		guildName = name
	}

	var template = `

**サーバー名**
%s

**実行コマンド**
%s

**開始日時**
%s
`

	now := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(template, guildName, strings.Join(command, " "), now)
	if err := SendSimpleEmbedMessage(s, AdminChannelID, "起動通知", msg); err != nil {
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
サーバー名: %s
`

	msg := fmt.Sprintf(template, guildName)
	if err := SendSimpleEmbedMessage(s, AdminChannelID, "停止コマンド通知", msg); err != nil {
		return CreateErr("停止コマンド実行通知メッセージを送信できません", err)
	}

	return nil
}
