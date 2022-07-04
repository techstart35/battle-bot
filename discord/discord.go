package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// ----- Discord -----
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
	session.AddHandler(battleHandler)

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
func battleHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content

	if cmd != Command {
		return
	}

	msg, err := sendEntryMessage(s, m)
	if err != nil {
		log.Println(err)
	}

	// TODO: テスト用に秒数を変更しています
	// 60秒後（残り60秒）にメッセージを送信
	time.Sleep(2 * time.Second)
	if err := sendCountDownMessage(s, msg, 60); err != nil {
		log.Println(err)
	}

	// 90秒後（残り30秒）にメッセージを送信
	time.Sleep(2 * time.Second)
	if err := sendCountDownMessage(s, msg, 30); err != nil {
		log.Println(err)
	}

	// 110秒後（残り10秒）にメッセージを送信
	time.Sleep(2 * time.Second)
	if err := sendCountDownMessage(s, msg, 10); err != nil {
		log.Println(err)
	}

	if err := sendStartMessage(s, msg); err != nil {
		log.Println(err)
	}
}
