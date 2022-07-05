package discord

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
)

// エントリーメッセージを送信します
func sendEntryMessage(s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.Message, error) {
	embedInfo := &discordgo.MessageEmbed{
		Title:       "⚔️ Giveaway Battle ⚔️",
		Description: "ねだるな！勝ち取れ🔥🔥",
		Color:       0x0099ff,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "主催者",
				Value:  fmt.Sprintf("<@%s>", m.Author.ID),
				Inline: false,
			},
			{
				Name:   "⚡️勝者",
				Value:  "1名",
				Inline: false,
			},
			{
				Name:   "⚡️エントリー",
				Value:  "⚔️のリアクション",
				Inline: false,
			},
			{
				Name:   "⚡️試合開始",
				Value:  "メッセージ送信から2分後",
				Inline: false,
			},
		},
	}

	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, embedInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	if err := s.MessageReactionAdd(m.ChannelID, msg.ID, "⚔️"); err != nil {
		return nil, errors.New(fmt.Sprintf("リアクションを付与できません: %v", err))
	}

	return msg, nil
}

// カウントダウンメッセージを送信します
func sendCountDownMessage(s *discordgo.Session, entryMsg *discordgo.Message, beforeStart uint) error {
	var color int
	switch beforeStart {
	case 60:
		color = 0x0099ff
	case 30:
		color = 0x3cb371
	case 10:
		color = 0xffd700
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("⚔️ Giveaway Battle開始まであと %d秒 ⚔️", beforeStart),
		Description: "Are You Ready?",
		Color:       color,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "エントリー",
				Value: fmt.Sprintf("[Jump!](https://discord.com/channels/%s/%s/%s)",
					os.Getenv("GUILD_ID"), entryMsg.ChannelID, entryMsg.ID),
				Inline: false,
			},
			{
				Name:   "中継先チャンネル",
				Value:  fmt.Sprintf("<#%s>", entryMsg.ChannelID),
				Inline: false,
			},
		},
	}

	_, err := s.ChannelMessageSendEmbed(entryMsg.ChannelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return nil
}

// 開始メッセージを送信します
func sendStartMessage(s *discordgo.Session, entryMsg *discordgo.Message) error {
	users, err := getReactedUsers(s, entryMsg)
	if err != nil {
		return errors.New(fmt.Sprintf("リアクションしたユーザーの取得に失敗しました: %v", err))
	}

	var tmpUser []string
	for _, v := range users {
		tmpUser = append(tmpUser, fmt.Sprintf("<@%s>", v.ID))
	}

	userStr := strings.Join(tmpUser, " ")

	embedInfo := &discordgo.MessageEmbed{
		Title:       "⚔️ Battle Start ⚔️",
		Description: fmt.Sprintf("挑戦者：%s", userStr),
		Color:       0xff0000,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "中継先チャンネル",
				Value:  fmt.Sprintf("<#%s>", entryMsg.ChannelID),
				Inline: true,
			},
		},
	}

	_, err = s.ChannelMessageSendEmbed(entryMsg.ChannelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return nil
}

// リアクションした人を取得します
func getReactedUsers(s *discordgo.Session, entryMsg *discordgo.Message) ([]*discordgo.User, error) {
	var users []*discordgo.User

	botName := os.Getenv("BOT_NAME")

	// 最大1000人まで参加可能（10 * 100）
	for i := 0; i < 10; i++ {
		var afterID string

		switch i {
		case 0:
			afterID = ""
		default:
			afterID = users[len(users)-1].ID
		}

		us, err := s.MessageReactions(entryMsg.ChannelID, entryMsg.ID, "⚔️", 100, "", afterID)
		if err != nil {
			return users, errors.New(fmt.Sprintf("リアクションをしたユーザーを取得できません: %v", err))
		}

		if len(us) == 1 && us[0].Username == botName {
			break
		}

		for _, u := range us {
			fmt.Println(i, u.Username, u.ID)
			if u.Username != botName {
				users = append(users, u)
			}
		}
	}

	return users, nil
}
