package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/handler/battle"
	"github.com/techstart35/battle-bot/handler/process"
	"github.com/techstart35/battle-bot/handler/reject_start"
	"github.com/techstart35/battle-bot/handler/stop"
)

// 文字入力のハンドラーを集約します
func TextHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	battle.BattleHandler(s, m)
	stop.StopHandler(s, m)
	process.ProcessHandler(s, m)
	reject_start.RejectStartHandler(s, m)
}
