package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/techstart35/battle-bot/handler/battle/normal"
	"github.com/techstart35/battle-bot/handler/process"
	"github.com/techstart35/battle-bot/handler/reject_start"
	"github.com/techstart35/battle-bot/handler/stop"
	"github.com/techstart35/battle-bot/shared/message"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const location = "Asia/Tokyo"

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(fmt.Sprintf(".envを読み込めません: %v", err))
	}
}

func main() {
	var Token = "Bot " + os.Getenv("APP_BOT_TOKEN")

	session, err := discordgo.New(Token)
	session.Token = Token
	if err != nil {
		message.SendErr(session, "discordのクライアントを作成できません", "none", "none", err)
		return
	}

	//イベントハンドラを追加
	session.AddHandler(normal.NormalBattleHandler)
	session.AddHandler(stop.StopHandler)
	session.AddHandler(process.ProcessHandler)
	session.AddHandler(reject_start.RejectStartHandler)

	if err = session.Open(); err != nil {
		message.SendErr(session, "discordを開けません", "none", "none", err)
		return
	}

	// 直近の関数（main）の最後に実行される
	defer func() {
		if err = session.Close(); err != nil {
			message.SendErr(session, "discordのクライアントを閉じれません", "none", "none", err)
		}
		return
	}()

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot
	return
}
