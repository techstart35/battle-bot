package message

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/guild"
	"strings"
	"time"
)

const AdminChannelID = "1003130506881277952"

// 開始時に自分のサーバーにメッセージを送信します
func SendStartMessageToAdmin(
	s *discordgo.Session,
	guildID, channelID string,
	command []string,
) error {
	guildName, err := guild.GetGuildName(s, guildID)
	if err != nil {
		return errors.NewError("ギルドを取得できません", err)
	}

	var template = `
⚔️｜サーバー名：**%s**
🔗｜起動チャンネル：%s
🚀｜実行コマンド：%s
`

	channelLink := shared.FormatChannelIDToLink(channelID)
	now := time.Now().Format("2006-01-02T15:04:05+09:00")
	msg := fmt.Sprintf(template, guildName, channelLink, strings.Join(command, " "))

	embedInfo := &discordgo.MessageEmbed{
		Title:       "Battle Royaleが起動されました",
		Description: msg,
		Color:       shared.ColorCyan,
		Timestamp:   now,
	}

	_, err = s.ChannelMessageSendEmbed(AdminChannelID, embedInfo)
	if err != nil {
		return errors.NewError("起動通知メッセージを送信できません", err)
	}

	return nil
}

// 正常終了時に自分のサーバーにメッセージを送信します
func SendNormalFinishMessageToAdmin(s *discordgo.Session, guildID string) error {
	guildName, err := guild.GetGuildName(s, guildID)
	if err != nil {
		return errors.NewError("ギルドを取得できません", err)
	}

	var template = `
✅️️｜サーバー名：**%s**
`

	now := time.Now().Format("2006-01-02T15:04:05+09:00")
	msg := fmt.Sprintf(template, guildName)

	embedInfo := &discordgo.MessageEmbed{
		Title:       "正常に終了しました",
		Description: msg,
		Color:       shared.ColorBlue,
		Timestamp:   now,
	}

	_, err = s.ChannelMessageSendEmbed(AdminChannelID, embedInfo)
	if err != nil {
		return errors.NewError("起動通知メッセージを送信できません", err)
	}

	return nil
}

// 停止コマンド実行時に自分のサーバーにメッセージを送信します
func SendStopMessageToAdmin(s *discordgo.Session, guildID string) error {
	guildName, err := guild.GetGuildName(s, guildID)
	if err != nil {
		return errors.NewError("ギルドを取得できません", err)
	}

	var template = `
**⚔️｜サーバー名**：%s
`

	now := time.Now().Format("2006-01-02T15:04:05+09:00")
	msg := fmt.Sprintf(template, guildName)

	embedInfo := &discordgo.MessageEmbed{
		Title:       "停止コマンドが実行されました",
		Description: msg,
		Color:       shared.ColorYellow,
		Timestamp:   now,
	}

	_, err = s.ChannelMessageSendEmbed(AdminChannelID, embedInfo)
	if err != nil {
		return errors.NewError("起動通知メッセージを送信できません", err)
	}

	return nil
}
