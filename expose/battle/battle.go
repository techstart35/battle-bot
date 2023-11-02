package battle

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/app/battle"
	"github.com/techstart35/battle-bot/app/list"
	"github.com/techstart35/battle-bot/app/reject_start"
	"github.com/techstart35/battle-bot/app/stop"
	"github.com/techstart35/battle-bot/gateway/di"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/message"
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
	case shared.Command().Start:
		battleApp := battle.NewBattleApp(ap)

		if err = battleApp.Battle(m.GuildID, m.ChannelID, m.Author.ID, input, 2); err != nil {
			req := message.SendErrReq{
				Message:   "バトルの実行に失敗しました",
				GuildID:   m.GuildID,
				ChannelID: m.ChannelID,
				Err:       err,
			}
			message.SendErr(s, req)
			return
		}
	case shared.Command().Start5Min:
		battleApp := battle.NewBattleApp(ap)

		if err = battleApp.Battle(m.GuildID, m.ChannelID, m.Author.ID, input, 5); err != nil {
			req := message.SendErrReq{
				Message:   "バトルの実行に失敗しました",
				GuildID:   m.GuildID,
				ChannelID: m.ChannelID,
				Err:       err,
			}
			message.SendErr(s, req)
			return
		}
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
	case shared.Command().List:
		listApp := list.NewList(ap)

		if err = listApp.List(); err != nil {
			req := message.SendErrReq{
				Message:   "リストの送信に失敗しました",
				GuildID:   m.GuildID,
				ChannelID: m.ChannelID,
				Err:       err,
			}
			message.SendErr(s, req)
			return
		}
	case shared.Command().RejectStart:
		rejectApp := reject_start.NewRejectStartApp(ap)

		if err = rejectApp.RejectStart(); err != nil {
			req := message.SendErrReq{
				Message:   "新規起動の停止に失敗しました",
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
