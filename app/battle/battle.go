package battle

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/app"
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/domain/model/battle"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/guild"
	"github.com/techstart35/battle-bot/shared/message"
	"strings"
)

// カスタムエラーです
var (
	startRejectedErr = fmt.Errorf("StartRejectedErr")
	isExistsErr      = fmt.Errorf("IsExistsErr")
	commandErr       = fmt.Errorf("CommandErr")
	isCanceledErr    = fmt.Errorf("IsCanceled")
)

// バトル構造体です
type BattleApp struct {
	*app.App
}

// バトル構造体を作成します
func NewBattleApp(app *app.App) *BattleApp {
	a := &BattleApp{}
	a.App = app

	return a
}

// バトルを実行します
//
// ユーザーへのメッセージはこの関数内でのみ記述します。
func (a *BattleApp) Battle(guildID, channelID, authorID string, input []string) error {
	gID, err := model.NewGuildID(guildID)
	if err != nil {
		return errors.NewError("ギルドIDを作成できません", err)
	}

	cID, err := model.NewChannelID(channelID)
	if err != nil {
		return errors.NewError("チャンネルIDを作成できません", err)
	}

	// 起動確認のバリデーションを行います
	switch err = a.validateEnabled(gID); err {
	case nil:
		break
	case startRejectedErr:
		if err = a.sendStartRejectedErrMsgToUser(cID); err != nil {
			return errors.NewError("起動停止済みメッセージを送信できません", err)
		}
		return nil
	case isExistsErr:
		if err = a.sendIsExistsErrToUser(cID); err != nil {
			return errors.NewError("重複起動エラーメッセージを送信できません", err)
		}
		return nil
	default:
		return errors.NewError("検証に失敗しました", err)
	}

	// コマンドのバリデーションを行います
	anChID, err := a.validateInput(input)
	switch err {
	case nil:
		break
	case commandErr:
		if err = a.sendCommandErrMsgToUser(cID); err != nil {
			return errors.NewError("コマンドエラーメッセージを送信できません", err)
		}
		return nil
	default:
		return errors.NewError("検証に失敗しました", err)
	}

	// battle構造体を作成します
	btl, err := battle.BuildBattle(guildID, channelID, anChID, authorID)
	if err != nil {
		return errors.NewError("battleを作成できません", err)
	}

	// Adminサーバーに起動メッセージを送信します
	if err = a.sendStartMsgToAdmin(gID, cID, input); err != nil {
		return errors.NewError("起動通知を送信できません")
	}

	// 永続化します
	{
		if err = a.Repo.Create(btl); err != nil {
			return errors.NewError("battleを作成できません", err)
		}

		// deferで発生したエラーのみ、直接Adminサーバーに送信します
		defer func() {
			if err = a.Repo.Delete(btl.GuildID()); err != nil {
				req := message.SendErrReq{
					Message:   "バトルを削除できません(defer)",
					GuildID:   btl.GuildID().String(),
					ChannelID: btl.ChannelID().String(),
					Err:       err,
				}
				message.SendErr(a.Session, req)
				return
			}
		}()
	}

	// エントリーメッセージを送信します
	if err = a.sendEntryMsgToUser(gID); err != nil {
		return errors.NewError("エントリーメッセージを送信できません", err)
	}

	// カウントダウンメッセージを送信します
	switch err = a.countDownScenario(gID); err {
	case nil:
		break
	case isCanceledErr:
		return nil
	default:
		return errors.NewError("カウントダウンメッセージを送信できません", err)
	}

	// 開始メッセージを送信します
	if err = a.sendStartMsgToUser(gID); err != nil {
		return errors.NewError("開始メッセージを送信できません", err)
	}

	// TODO: バトルメッセージを送信

	// 正常終了通知を送信します
	//
	// [Notice] メソッドの一番最後に実行します
	if err = a.sendFinishedMsgToAdmin(gID, cID); err != nil {
		return errors.NewError("正常終了通知を送信できません", err)
	}

	return nil
}

// 起動可能か検証します
//
// コールする側で startRejectErr / isExistsErr のエラーハンドリングを行います。
func (a *BattleApp) validateEnabled(guildID model.GuildID) error {
	// 新規の起動が停止されているかを確認します
	if a.Query.IsStartRejected() {
		return startRejectedErr
	}

	// 既に起動しているバトルがあるか確認します
	btls, err := a.Query.FindAll()
	switch err {
	case nil:
		for _, btl := range btls {
			if btl.GuildID().Equal(guildID) {
				return isExistsErr
			}
		}
	case errors.NotFoundErr:
		return nil
	default:
		return errors.NewError("全てのバトルを取得できません", err)
	}

	return nil
}

// 引数の確認をします
//
// コールする側で commandErr のエラーハンドリングを行います。
func (a *BattleApp) validateInput(input []string) (string, error) {
	if len(input) > 1 {
		ti := strings.TrimLeft(input[1], "<#")
		anotherChannelID := strings.TrimRight(ti, ">")

		// 配信チャンネルのチャンネルIDが正しいことを検証
		if _, err := a.Session.Channel(anotherChannelID); err != nil {
			return "", commandErr
		}

		return anotherChannelID, nil
	}

	return "", nil
}

// startRejectedErr のエラーメッセージを送信します
func (a *BattleApp) sendStartRejectedErrMsgToUser(channelID model.ChannelID) error {
	const MessageTmpl = `
メンテナンスのため、botの起動を一時停止しております。
数分後に再度お試しください。
`

	embedInfo := &discordgo.MessageEmbed{
		Title:       "INFO",
		Description: MessageTmpl,
		Color:       shared.ColorBlack,
		Timestamp:   shared.GetNowTimeStamp(),
	}

	_, err := a.Session.ChannelMessageSendEmbed(channelID.String(), embedInfo)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}

// IsExistsErr のメッセージを送信します
func (a *BattleApp) sendIsExistsErrToUser(channelID model.ChannelID) error {
	const MessageTmpl = `
このサーバーで起動しているbattleが存在します。

キャンセル済みの場合は反映までお待ちください。
（最大1分かかります）
`

	embedInfo := &discordgo.MessageEmbed{
		Title:       "INFO",
		Description: MessageTmpl,
		Color:       shared.ColorBlack,
		Timestamp:   shared.GetNowTimeStamp(),
	}

	_, err := a.Session.ChannelMessageSendEmbed(channelID.String(), embedInfo)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}

// CommandErr のメッセージを送信します
func (a *BattleApp) sendCommandErrMsgToUser(channelID model.ChannelID) error {
	const MessageTmpl = `
コマンドが間違っているか、チャンネルの権限が不足しています。
`

	embedInfo := &discordgo.MessageEmbed{
		Title:       "ERROR",
		Description: MessageTmpl,
		Color:       shared.ColorBlack,
		Timestamp:   shared.GetNowTimeStamp(),
	}

	_, err := a.Session.ChannelMessageSendEmbed(channelID.String(), embedInfo)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}

// Adminサーバーに起動通知を送信します
func (a *BattleApp) sendStartMsgToAdmin(
	guildID model.GuildID,
	channelID model.ChannelID,
	input []string,
) error {
	var MessageTmpl = `
⚔️｜サーバー名：**%s**
🔗｜起動チャンネル：%s
🚀｜実行コマンド：%s
`

	guildName, err := guild.GetGuildName(a.Session, guildID.String())
	if err != nil {
		return errors.NewError("ギルド名を取得できません", err)
	}

	msg := fmt.Sprintf(
		MessageTmpl,
		guildName,
		shared.FormatChannelIDToLink(channelID.String()),
		strings.Join(input, " "),
	)

	embedInfo := &discordgo.MessageEmbed{
		Title:       "Battle Royaleが起動されました",
		Description: msg,
		Color:       shared.ColorCyan,
		Timestamp:   shared.GetNowTimeStamp(),
	}

	_, err = a.Session.ChannelMessageSendEmbed(message.AdminChannelID, embedInfo)
	if err != nil {
		return errors.NewError("起動通知メッセージを送信できません", err)
	}

	return nil
}

// 正常終了時にAdminサーバーに通知します
func (a *BattleApp) sendFinishedMsgToAdmin(
	guildID model.GuildID,
	channelID model.ChannelID,
) error {
	var MessageTmpl = `
✅️️｜サーバー名：**%s**
`

	guildName, err := guild.GetGuildName(a.Session, guildID.String())
	if err != nil {
		return errors.NewError("ギルド名を取得できません", err)
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       "正常に終了しました",
		Description: fmt.Sprintf(MessageTmpl, guildName),
		Color:       shared.ColorBlue,
		Timestamp:   shared.GetNowTimeStamp(),
	}

	_, err = a.Session.ChannelMessageSendEmbed(message.AdminChannelID, embedInfo)
	if err != nil {
		return errors.NewError("起動通知メッセージを送信できません", err)
	}

	return nil
}
