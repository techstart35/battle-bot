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
⚔️｜サーバー名：**%s**
🔗｜起動チャンネル：%s
✅｜実行コマンド：%s
`

	channelLink := FormatChannelIDToLink(channelID)
	now := time.Now().Format("2006-01-02T15:04:05+09:00")
	msg := fmt.Sprintf(template, guildName, channelLink, strings.Join(command, " "))

	embedInfo := &discordgo.MessageEmbed{
		Title:       "Battle Royaleが起動されました",
		Description: msg,
		Color:       ColorCyan,
		Timestamp:   now,
	}

	_, err := s.ChannelMessageSendEmbed(AdminChannelID, embedInfo)
	if err != nil {
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
**⚔️｜サーバー名**：%s
`

	now := time.Now().Format("2006-01-02T15:04:05+09:00")
	msg := fmt.Sprintf(template, guildName)

	embedInfo := &discordgo.MessageEmbed{
		Title:       "停止コマンド通知",
		Description: msg,
		Color:       ColorYellow,
		Timestamp:   now,
	}

	_, err := s.ChannelMessageSendEmbed(AdminChannelID, embedInfo)
	if err != nil {
		return CreateErr("起動通知メッセージを送信できません", err)
	}

	return nil
}
