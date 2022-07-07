package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/message"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	Command = "b"
)

// Discordでメッセージを送信します
func SendDiscord() {
	var Token = "Bot " + os.Getenv("APP_BOT_TOKEN")

	session, err := discordgo.New(Token)
	session.Token = Token
	if err != nil {
		log.Printf(fmt.Sprintf("discordのクライアントを作成できません: %v", err))
	}

	//イベントハンドラを追加
	session.AddHandler(BattleHandler)

	if err := session.Open(); err != nil {
		log.Printf(fmt.Sprintf("discordを開けません: %v", err))
	}

	// 直近の関数（main）の最後に実行される
	defer func() {
		if err := session.Close(); err != nil {
			log.Printf(fmt.Sprintf("discordののクライアントを閉じれません: %v", err))
		}
	}()

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot
	return
}

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

	msg, err := message.SendEntryMessage(s, m, anotherChannelID)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(60 * time.Second)

	// 60秒後（残り60秒）にメッセージを送信
	if err := message.SendCountDownMessage(s, msg, 60, anotherChannelID); err != nil {
		log.Println(err)
	}

	time.Sleep(30 * time.Second)

	// 残り30秒アナウンス
	if err := message.SendCountDownMessage(s, msg, 30, anotherChannelID); err != nil {
		log.Println(err)
	}

	time.Sleep(20 * time.Second)

	// 残り10秒アナウンス
	if err := message.SendCountDownMessage(s, msg, 10, anotherChannelID); err != nil {
		log.Println(err)
	}

	time.Sleep(10 * time.Second)

	usrs, err := message.SendStartMessage(s, msg, anotherChannelID)
	if err != nil {
		log.Println(err)
	}

	time.Sleep(10 * time.Second)

	if err := message.BattleMessageHandler(s, usrs, msg, anotherChannelID); err != nil {
		log.Println(err)
	}
}
