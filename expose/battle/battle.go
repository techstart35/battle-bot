package battle

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/app/stop"
	"github.com/techstart35/battle-bot/gateway/di"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/message"
	"strings"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	ap, err := di.InitApp(s)
	if err != nil {
		req := message.SendErrReq{
			Message:   "アプリケーションの初期化に失敗しました",
			GuildID:   m.GuildID,
			ChannelID: m.ChannelID,
			Err:       err,
		}
		message.SendErr(s, req)
		return
	}

	input := strings.Split(m.Content, " ")
	cmd := input[0]

	switch cmd {
	case shared.Command().Stop:
		stopApp := stop.NewStopApp(ap)

		if err = stopApp.StopBattle(m.GuildID, m.ChannelID); err != nil {
			req := message.SendErrReq{
				Message:   "バトルの停止に失敗しました",
				GuildID:   m.GuildID,
				ChannelID: m.ChannelID,
				Err:       err,
			}
			message.SendErr(s, req)
			return
		}
	}

	return
}
