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
				Name:   "▼主催者",
				Value:  fmt.Sprintf("<@%s>", m.Author.ID),
				Inline: false,
			},
			{
				Name:   "▼勝者",
				Value:  "1名",
				Inline: false,
			},
			{
				Name:   "▼エントリー",
				Value:  "⚔️のリアクション",
				Inline: false,
			},
			{
				Name:   "▼試合開始",
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
				Name: "▼エントリー",
				Value: fmt.Sprintf("[Jump!](https://discord.com/channels/%s/%s/%s)",
					os.Getenv("GUILD_ID"), entryMsg.ChannelID, entryMsg.ID),
				Inline: false,
			},
			{
				Name:   "▼このチャンネルにも戦いの様子を送信します",
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
func sendStartMessage(s *discordgo.Session, entryMsg *discordgo.Message) ([]*discordgo.User, error) {
	users, err := getReactedUsers(s, entryMsg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("リアクションしたユーザーの取得に失敗しました: %v", err))
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
				Name:   "▼このチャンネルにも戦いの様子を送信します",
				Value:  fmt.Sprintf("<#%s>", entryMsg.ChannelID),
				Inline: true,
			},
		},
	}

	_, err = s.ChannelMessageSendEmbed(entryMsg.ChannelID, embedInfo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return users, nil
}

// NoEntryのメッセージを送信します
func sendNoEntryMessage(s *discordgo.Session, entryMessage *discordgo.Message) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       "No Entry",
		Description: "エントリーが無かったため、試合は開始されません",
		Color:       0xff0000,
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return nil
}

// Battleのメッセージを送信します
func sendBattleMessage(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	description string,
	round int,
) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("第%d回戦", round),
		Description: description,
		Color:       0xff0000,
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return nil
}

// Winnerのメッセージを送信します
func sendWinnerMessage(s *discordgo.Session, entryMessage *discordgo.Message, winner *discordgo.User) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       "👑 Winner 👑",
		Description: fmt.Sprintf("勝者：<@%s>", winner.ID),
		Color:       0xff0000,
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	msg, err := s.ChannelMessageSend(entryMessage.ChannelID, fmt.Sprintf("<@%s>さん、おめでとう🎉", winner.ID))
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	if err := s.MessageReactionAdd(msg.ChannelID, msg.ID, "🎉"); err != nil {
		return errors.New(fmt.Sprintf("リアクションを付与できません: %v", err))
	}

	return nil
}
