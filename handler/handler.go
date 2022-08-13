package handler

import "github.com/bwmarrin/discordgo"

// 文字入力のハンドラーを集約します
func TextHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	BattleHandler(s, m)
	StopHandler(s, m)
	ProcessHandler(s, m)
	RejectStartHandler(s, m)
}
