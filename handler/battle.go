package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/handler/message/battle"
	"github.com/techstart35/battle-bot/handler/message/countdown"
	"github.com/techstart35/battle-bot/handler/message/entry"
	"github.com/techstart35/battle-bot/handler/message/start"
	"github.com/techstart35/battle-bot/shared"
	"strings"
	"time"
)

// Battleを実行します
//
// 1: b
// 2: b <#channelID>
func BattleHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	input := strings.Split(m.Content, " ")

	cmd := input[0]
	if cmd != shared.Command().Start {
		return
	}

	// 新規の起動が停止されているかを確認します
	if shared.IsStartRejected {
		if err := shared.SendSimpleEmbedMessage(
			s, m.ChannelID, "Info", "メンテナンスのため、botの起動を一時停止しております。数分後に再度お試しください。",
		); err != nil {
			shared.SendErr(s, "RejectStartメッセージを送信できません", m.GuildID, m.ChannelID, err)
			return
		}

		return
	}

	var anotherChannelID string

	if len(input) >= 2 {
		// チャンネルIDを設定
		anotherChannelID = strings.TrimLeft(input[1], "<#")
		anotherChannelID = strings.TrimRight(anotherChannelID, ">")

		// チャンネルIDが正しいことを検証
		if _, err := s.Channel(anotherChannelID); err != nil {
			if err := shared.SendSimpleEmbedMessage(s, m.ChannelID, "ERROR", "コマンドが正しくありません"); err != nil {
				shared.SendErr(s, "コマンドエラーメッセージを送信できません", m.GuildID, m.ChannelID, err)
				return
			}
			return
		}
	}

	// チャンネル一覧に追加
	shared.IsProcessing[m.ChannelID] = true

	msg, err := entry.SendEntryMessage(s, m, anotherChannelID)
	if err != nil {
		shared.SendErr(s, "エントリーメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	time.Sleep(60 * time.Second)

	// 60秒後（残り60秒）にメッセージを送信
	if err := countdown.SendCountDownMessage(s, msg, 60, anotherChannelID); err != nil {
		shared.SendErr(s, "60秒前カウントダウンメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	time.Sleep(30 * time.Second)

	// 残り30秒アナウンス
	if err := countdown.SendCountDownMessage(s, msg, 30, anotherChannelID); err != nil {
		shared.SendErr(s, "30秒前カウントダウンメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	time.Sleep(20 * time.Second)

	// 残り10秒アナウンス
	if err := countdown.SendCountDownMessage(s, msg, 10, anotherChannelID); err != nil {
		shared.SendErr(s, "10秒前カウントダウンメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	time.Sleep(10 * time.Second)

	// 開始メッセージ
	usrs, err := start.SendStartMessage(s, msg, anotherChannelID)
	if err != nil {
		shared.SendErr(s, "開始メッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	time.Sleep(10 * time.Second)

	// バトルメッセージ
	if err := battle.BattleMessageHandler(s, usrs, msg, anotherChannelID); err != nil {
		shared.SendErr(s, "バトルメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	// チャンネル一覧から削除
	delete(shared.IsProcessing, m.ChannelID)
}
