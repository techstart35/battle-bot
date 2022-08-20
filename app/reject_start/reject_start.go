package reject_start

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/app"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/message"
)

// 新規起動停止の構造体です
type RejectStartApp struct {
	*app.App
}

// 新規起動停止の構造体を作成します
func NewRejectStartApp(app *app.App) *RejectStartApp {
	a := &RejectStartApp{}
	a.App = app

	return a
}

// 新規起動を停止します
func (a *RejectStartApp) RejectStart() error {
	// 起動停止済かを確認
	if a.Query.IsStartRejected() {
		if err := a.sendAlreadyRejectedMsgToAdmin(); err != nil {
			return errors.NewError("起動停止済メッセージを送信できません", err)
		}
		return nil
	}

	a.Repo.RejectStart()

	// 停止完了メッセージを送信
	if err := a.sendStartRejectedMsgToAdmin(); err != nil {
		return errors.NewError("起動停止完了メッセージを送信できません", err)
	}

	return nil
}

// Adminサーバーに起動停止済メッセージを送信します
func (a *RejectStartApp) sendAlreadyRejectedMsgToAdmin() error {
	const MessageTmpl = "すでに新規起動が停止されています"

	embedInfo := &discordgo.MessageEmbed{
		Description: MessageTmpl,
		Color:       shared.ColorBlack,
	}

	_, err := a.Session.ChannelMessageSendEmbed(message.AdminChannelID, embedInfo)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}

// Adminサーバーに起動停止完了メッセージを送信します
func (a *RejectStartApp) sendStartRejectedMsgToAdmin() error {
	const MessageTmpl = "新規起動を停止しました"

	embedInfo := &discordgo.MessageEmbed{
		Title:       "新規起動の停止",
		Description: MessageTmpl,
		Color:       shared.ColorPink,
	}

	_, err := a.Session.ChannelMessageSendEmbed(message.AdminChannelID, embedInfo)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}
