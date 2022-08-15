package normal

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/handler/battle"
	battleMessage "github.com/techstart35/battle-bot/handler/battle/message/battle/scenario/normal"
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

	// 起動可能な状態か確認します
	ok, err := battle.CheckBeforeStartAndSendMessage(s, m.GuildID, m.ChannelID)
	if err != nil {
		message.SendErr(s, "起動前のチェックができません", m.GuildID, m.ChannelID, err)
		return
	}
	if !ok {
		return
	}

	var anotherChannelID string

	// 引数の確認をします
	anChID, err := CheckInput(s, m.ChannelID, input)
	if err != nil {
		message.SendErr(s, "引数チェックに失敗しました", m.GuildID, m.ChannelID, err)
		return
	}
	anotherChannelID = anChID

	// Adminサーバーに起動メッセージを送信します
	//
	// Notice: ここでエラーが発生しても処理は継続させます
	if err := message.SendStartMessageToAdmin(s, m.GuildID, m.ChannelID, input); err != nil {
		message.SendErr(s, "起動通知をAdminサーバーに送信できません", m.GuildID, m.ChannelID, err)
	}

	// チャンネル一覧に追加します
	shared.SetNewProcess(m.GuildID)
	defer shared.DeleteProcess(m.GuildID)

	// エントリーメッセージを送信します
	msg, err := entry.SendEntryMessage(s, m, anotherChannelID)
	if err != nil {
		message.SendErr(s, "エントリーメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	// カウントダウンメッセージを送信します
	ok, err = countdown.CountDownScenario(s, msg, anotherChannelID)
	if err != nil {
		message.SendErr(s, "カウントダウンメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}
	if !ok {
		return
	}

	// 開始メッセージを送信します
	usrs, err := start.SendStartMessage(s, msg, anotherChannelID)
	if err != nil {
		message.SendErr(s, "開始メッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	if battle.IsCanceledCheckAndSleep(10, m.GuildID) {
		return
	}

	// バトルメッセージを送信します
	if err = battleMessage.NormalBattleMessageScenario(s, usrs, msg, anotherChannelID); err != nil {
		message.SendErr(s, "バトルメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	// 正常終了のメッセージを送信します
	if err = message.SendNormalFinishMessageToAdmin(s, m.GuildID); err != nil {
		message.SendErr(s, "終了通知をAdminサーバーに送信できません", m.GuildID, m.ChannelID, err)
		return
	}
}
