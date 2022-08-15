package normal

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/handler/battle"
	battleMessage "github.com/techstart35/battle-bot/handler/battle/message/battle"
	"github.com/techstart35/battle-bot/handler/battle/message/countdown"
	"github.com/techstart35/battle-bot/handler/battle/message/entry"
	"github.com/techstart35/battle-bot/handler/battle/message/start"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/message"
	"strings"
)

// Battleを実行します
//
// 1: b
// 2: b <#channelID>
func NormalBattleHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	input := strings.Split(m.Content, " ")

	cmd := input[0]
	// コマンドが一致するか確認します
	if cmd != shared.Command().Start {
		return
	}

	// 新規の起動が停止されているかを確認します
	if shared.IsStartRejected {
		if err := message.SendSimpleEmbedMessage(
			s, m.ChannelID, "INFO", "メンテナンスのため、botの起動を一時停止しております。数分後に再度お試しください。", 0,
		); err != nil {
			message.SendErr(s, "RejectStartメッセージを送信できません", m.GuildID, m.ChannelID, err)
			return
		}

		return
	}

	// すでに起動しているbattleを確認します
	if shared.IsProcessExists(m.GuildID) {
		msg := `
このサーバーで起動しているbattleが存在します。

キャンセル済みの場合は反映までお待ちください。
（最大1分かかります）
`
		if err := message.SendSimpleEmbedMessage(s, m.ChannelID, "INFO", msg, 0); err != nil {
			message.SendErr(s, "RejectStartメッセージを送信できません", m.GuildID, m.ChannelID, err)
			return
		}

		return
	}

	var anotherChannelID string

	if len(input) >= 2 {
		// チャンネルIDを設定
		anotherChannelID = strings.TrimLeft(input[1], "<#")
		anotherChannelID = strings.TrimRight(anotherChannelID, ">")

		// チャンネルIDが正しいことを検証
		if _, err := s.Channel(anotherChannelID); err != nil {
			msg := "コマンドが間違っているか、チャンネルの権限が不足しています。"
			if err := message.SendSimpleEmbedMessage(s, m.ChannelID, "ERROR", msg, 0); err != nil {
				message.SendErr(s, "コマンドエラーメッセージを送信できません", m.GuildID, m.ChannelID, err)
				return
			}
			return
		}
	}

	// チャンネル一覧から削除
	defer shared.DeleteProcess(m.GuildID)

	// Adminサーバーに起動メッセージを送信します
	//
	// Notice: ここでエラーが発生しても処理は継続させます
	if err := message.SendStartMessageToAdmin(s, m.GuildID, m.ChannelID, input); err != nil {
		message.SendErr(s, "起動通知をAdminサーバーに送信できません", m.GuildID, m.ChannelID, err)
	}

	// チャンネル一覧に追加
	shared.SetNewProcess(m.GuildID)

	msg, err := entry.SendEntryMessage(s, m, anotherChannelID)
	if err != nil {
		message.SendErr(s, "エントリーメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	if battle.IsCanceledCheckAndSleep(60, m.GuildID) {
		return
	}

	// 60秒後（残り60秒）にメッセージを送信
	if err = countdown.SendCountDownMessage(s, msg, 60, anotherChannelID); err != nil {
		message.SendErr(s, "60秒前カウントダウンメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	if battle.IsCanceledCheckAndSleep(30, m.GuildID) {
		return
	}

	// 残り30秒アナウンス
	if err = countdown.SendCountDownMessage(s, msg, 30, anotherChannelID); err != nil {
		message.SendErr(s, "30秒前カウントダウンメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	if battle.IsCanceledCheckAndSleep(20, m.GuildID) {
		return
	}

	// 残り10秒アナウンス
	if err = countdown.SendCountDownMessage(s, msg, 10, anotherChannelID); err != nil {
		message.SendErr(s, "10秒前カウントダウンメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	if battle.IsCanceledCheckAndSleep(10, m.GuildID) {
		return
	}

	// 開始メッセージ
	usrs, err := start.SendStartMessage(s, msg, anotherChannelID)
	if err != nil {
		message.SendErr(s, "開始メッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	if battle.IsCanceledCheckAndSleep(10, m.GuildID) {
		return
	}

	// バトルメッセージ
	if err = battleMessage.BattleMessageScenario(s, usrs, msg, anotherChannelID); err != nil {
		message.SendErr(s, "バトルメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	// 正常終了のメッセージを送信します
	if err = message.SendNormalFinishMessageToAdmin(s, m.GuildID); err != nil {
		message.SendErr(s, "終了通知をAdminサーバーに送信できません", m.GuildID, m.ChannelID, err)
		return
	}
}