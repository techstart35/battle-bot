package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/techstart35/battle-bot/discord"
	"log"
	"time"
)

const location = "Asia/Tokyo"

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	loadEnv()
	discord.SendDiscord()
}

// .envファイルを読み込みます
func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(fmt.Sprintf(".envを読み込めません: %v", err))
	}
}
