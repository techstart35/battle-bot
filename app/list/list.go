package list

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/app"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/guild"
	"github.com/techstart35/battle-bot/shared/message"
	"strings"
)

// 一覧構造体です
type ListApp struct {
	*app.App
}

// 一覧構造体を作成します
func NewList(app *app.App) *ListApp {
	a := &ListApp{}
	a.App = app

	return a
}

// 一覧をAdminサーバーに送信します
func (a *ListApp) List() error {
	var msg string

	btls, err := a.Query.FindAll()
	switch err {
	// 正常な場合: ステータスを付与して送信
	case nil:
		m := make([]string, 0)
		for _, btl := range btls {
			status := "✅｜起動中"
			if btl.IsCanceled() {
				status = "🌙｜キャンセル済"
			}

			guildName, err := guild.GetGuildName(a.Session, btl.GuildID().String())
			if err != nil {
				return errors.NewError("一覧を送信できません")
			}

			m = append(m, fmt.Sprintf("%s｜サーバー名: **%s**", status, guildName))
		}

		msg = strings.Join(m, "\n")

	// 起動プロセスが無い場合: メッセージを送信
	case errors.NotFoundErr:
		msg = "起動中のバトルがありません"

	// エラーが発生した場合: エラーハンドリング
	default:
		return errors.NewError("実行中のプロセスの取得に失敗しました", err)
	}

	req := &discordgo.MessageEmbed{
		Title:       "一覧の表示",
		Description: msg,
		Color:       shared.ColorPink,
	}
	_, err = a.Session.ChannelMessageSendEmbed(message.AdminChannelID, req)
	if err != nil {
		return errors.NewError("一覧をAdminサーバーに送信できません", err)
	}

	return nil
}
