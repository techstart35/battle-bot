package battle

// ---------------------------
// Battleメッセージに関する共通処理を記述します
// ---------------------------

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/message"
	"time"
)

// キャンセル確認とSleep処理を実行します
func IsCanceledCheckAndSleep(second int, guildID string) bool {
	if shared.IsCanceled(guildID) {
		return true
	}

	time.Sleep(time.Duration(second) * time.Second)

	if shared.IsCanceled(guildID) {
		return true
	}

	return false
}

// 新規起動が停止していた場合のメッセージテンプレートです
const rejectStartTmpl = `
メンテナンスのため、botの起動を一時停止しております。
数分後に再度お試しください。
`

// プロセスが存在していた場合のメッセージテンプレートです
const processExistsTmpl = `
このサーバーで起動しているbattleが存在します。

キャンセル済みの場合は反映までお待ちください。
（最大1分かかります）
`

// 起動前のチェックを行います
//
// 起動停止の場合のみfalseを返します。
//
// 起動停止の場合は、この関数内でINFOメッセージを送信します。
func CheckBeforeStartAndSendMessage(
	s *discordgo.Session,
	guildID, channelID string,
) (bool, error) {
	// 新規の起動が停止されているかを確認します
	if shared.IsStartRejected {
		if err := message.SendSimpleEmbedMessage(s, channelID, "INFO", rejectStartTmpl, 0); err != nil {
			return false, errors.NewError("RejectStartメッセージを送信できません", err)
		}

		return false, nil
	}

	// すでに起動しているbattleを確認します
	if shared.IsProcessExists(guildID) {
		if err := message.SendSimpleEmbedMessage(s, channelID, "INFO", processExistsTmpl, 0); err != nil {
			return false, errors.NewError("RejectStartメッセージを送信できません")
		}

		return false, nil
	}

	return true, nil
}
