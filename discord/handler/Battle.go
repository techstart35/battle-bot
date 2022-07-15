package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/message"
	"github.com/techstart35/battle-bot/discord/message/battle"
	"github.com/techstart35/battle-bot/discord/message/countdown"
	"github.com/techstart35/battle-bot/discord/message/entry"
	"github.com/techstart35/battle-bot/discord/message/start"
	"github.com/techstart35/battle-bot/discord/shared"
	"log"
	"strings"
	"time"
)

// battleを実行します
func BattleHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	input := strings.Split(m.Content, " ")
	cmd := input[0]

	if cmd != Command {
		return
	}

	var anotherChannelID string

	if len(input) >= 2 {
		// チャンネルIDを設定
		anotherChannelID = strings.TrimLeft(input[1], "<#")
		anotherChannelID = strings.TrimRight(anotherChannelID, ">")

		// チャンネルIDが正しいことを検証
		if _, err := s.Channel(anotherChannelID); err != nil {
			if err := message.SendSimpleEmbedMessage(s, m.ChannelID, "ERROR", "コマンドが正しくありません"); err != nil {
				log.Println(err)
			}
			return
		}
	}

	// チャンネル一覧に追加
	shared.IsProcessing[m.ChannelID] = true

	msg, err := entry.SendEntryMessage(s, m, anotherChannelID)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(60 * time.Second)

	// 60秒後（残り60秒）にメッセージを送信
	if err := countdown.SendCountDownMessage(s, msg, 60, anotherChannelID); err != nil {
		log.Println(err)
	}

	time.Sleep(30 * time.Second)

	// 残り30秒アナウンス
	if err := countdown.SendCountDownMessage(s, msg, 30, anotherChannelID); err != nil {
		log.Println(err)
	}

	time.Sleep(20 * time.Second)

	// 残り10秒アナウンス
	if err := countdown.SendCountDownMessage(s, msg, 10, anotherChannelID); err != nil {
		log.Println(err)
	}

	time.Sleep(10 * time.Second)

	usrs, err := start.SendStartMessage(s, msg, anotherChannelID)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(10 * time.Second)

	if err := battle.BattleMessageHandler(s, usrs, msg, anotherChannelID); err != nil {
		log.Println(err)
	}

	// チャンネル一覧から削除
	delete(shared.IsProcessing, m.ChannelID)
}
