package stop

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/app"
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/guild"
	"github.com/techstart35/battle-bot/shared/message"
)

// 停止構造体です
type StopApp struct {
	*app.App
}

// 停止構造体を作成します
func NewStopApp(app *app.App) *StopApp {
	a := &StopApp{}
	a.App = app

	return a
}

// 停止処理を実行します
//
// ユーザーへのメッセージはこの関数内でのみ記述します。
func (a *StopApp) StopBattle(guildID, channelID string) error {
	gID, err := model.NewGuildID(guildID)
	if err != nil {
		return errors.NewError("ギルドIDを作成できません", err)
	}

	cID, err := model.NewChannelID(channelID)
	if err != nil {
		return errors.NewError("チャンネルIDを作成できません", err)
	}

	// Adminに停止コマンド起動メッセージを送信
	if err := a.sendStopMsgToAdmin(gID, cID); err != nil {
		return errors.NewError("Adminに停止コマンド起動メッセージを送信できません", err)
	}

	btl, err := a.Query.FindByGuildID(gID)
	if err != nil && err != errors.NotFoundErr {
		return errors.NewError("ギルドIDでバトルを取得できません", err)
	}

	// バリデーションを行います
	{
		if err == errors.NotFoundErr || btl.IsCanceled() {
			// ユーザーに停止不可メッセージを送信します
			if err = a.sendValidateErrMsgToUser(a.Session, cID); err != nil {
				return errors.NewError("ユーザーに停止不可メッセージを送信できません", err)
			}

			return nil
		}
	}

	// 停止処理を実行
	btl.Cancel()
	if err = a.Repo.Update(btl); err != nil {
		return errors.NewError("更新に失敗しました", err)
	}

	// 停止完了メッセージを送信
	if err = a.sendStoppedMsgToUser(a.Session, btl.ChannelID()); err != nil {
		return errors.NewError("停止完了メッセージを送信できません", err)
	}

	return nil
}

// 停止処理完了メッセージをユーザーに送信します
func (a *StopApp) sendStoppedMsgToUser(s *discordgo.Session, cID model.ChannelID) error {
	const MessageTmpl = `
このサーバーで起動されたバトルをキャンセルしました。
（反映まで最大1分かかります）
`

	req := &discordgo.MessageEmbed{
		Title:       "キャンセル処理の実行",
		Description: MessageTmpl,
		Color:       shared.ColorBlack,
	}
	_, err := s.ChannelMessageSendEmbed(cID.String(), req)
	if err != nil {
		return errors.NewError("停止処理完了メッセージをユーザーに送信できません", err)
	}

	return nil
}

// バリデーションエラーメッセージをユーザーに送信します
func (a *StopApp) sendValidateErrMsgToUser(s *discordgo.Session, cID model.ChannelID) error {
	const MessageTmpl = `
このサーバーで起動されたバトルが無いか、
キャンセル済みとなっています。
`

	embedInfo := &discordgo.MessageEmbed{
		Title:       "ERROR",
		Description: MessageTmpl,
		Color:       shared.ColorBlack,
	}
	_, err := s.ChannelMessageSendEmbed(cID.String(), embedInfo)
	if err != nil {
		return errors.NewError("バリデーションエラーメッセージをユーザーに送信できません", err)
	}

	return nil
}

// 停止処理の起動をAdminに送信します
//
// [注意]バトルを取得できない可能性もあるため、引数のIDはコマンド実行時のIDを入れます。
func (a *StopApp) sendStopMsgToAdmin(
	guildID model.GuildID,
	channelID model.ChannelID,
) error {
	const MessageTmpl = `
⚔️｜サーバー名：**%s**
🔗｜チャンネル：**%s**
`

	guildName, err := guild.GetGuildName(a.Session, guildID.String())
	if err != nil {
		return errors.NewError("ギルドを取得できません", err)
	}

	embedInfo := &discordgo.MessageEmbed{
		Title: "停止コマンドが実行されました",
		Description: fmt.Sprintf(
			MessageTmpl,
			guildName,
			shared.FormatChannelIDToLink(channelID.String()),
		),
		Color:     shared.ColorYellow,
		Timestamp: shared.GetNowTimeStamp(),
	}

	_, err = a.Session.ChannelMessageSendEmbed(message.AdminChannelID, embedInfo)
	if err != nil {
		return errors.NewError("起動通知メッセージを送信できません", err)
	}

	return nil
}
