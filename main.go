package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/techstart35/battle-bot/expose/battle"
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
		req := message.SendErrReq{
			Message:   "discordのクライアントを作成できません",
			GuildID:   "none",
			ChannelID: "none",
			Err:       err,
		}
		message.SendErr(session, req)
		return
	}

	//イベントハンドラを追加
	session.AddHandler(battle.Handler)

	if err = session.Open(); err != nil {
		req := message.SendErrReq{
			Message:   "discordを開けません",
			GuildID:   "none",
			ChannelID: "none",
			Err:       err,
		}
		message.SendErr(session, req)
		return
	}

	// 直近の関数（main）の最後に実行される
	defer func() {
		if err = session.Close(); err != nil {
			req := message.SendErrReq{
				Message:   "discordのクライアントを閉じれません",
				GuildID:   "none",
				ChannelID: "none",
				Err:       err,
			}
			message.SendErr(session, req)
		}
		return
	}()

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stopBot
	return
}
