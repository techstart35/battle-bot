package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	session.AddHandler(handler.BattleHandler)
	session.AddHandler(handler.StopHandler)
	session.AddHandler(handler.ProcessHandler)
	session.AddHandler(handler.RejectStartHandler)

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
