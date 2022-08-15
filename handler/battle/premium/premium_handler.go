package premium

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
// channelID(#判定) > userID(@判定) > winner(数字判定)
//
// 1: pb
// 2: pb <#channelID>
// 3: pb <@userID>
// 4: pb 3w
// 5: pb <#channelID> <@userID> 3w
func PremiumBattleHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	input := strings.Split(m.Content, " ")

	cmd := input[0]
	// コマンドが一致するか確認します
	if cmd != shared.Command().Start {
		return
	}

	ok, err := battle.CheckBeforeStartAndSendMessage(s, m.GuildID, m.ChannelID)
	if err != nil {
		message.SendErr(s, "起動前のチェックができません", m.GuildID, m.ChannelID, err)
		return
	}
	if !ok {
		return
	}

	args, err := CheckInput(s, input)
	if err != nil {
		message.SendErr(s, "コマンドが正しくありません", m.GuildID, m.ChannelID, err)
		return
	}

	// Adminサーバーに起動メッセージを送信します
	//
	// Notice: ここでエラーが発生しても処理は継続させます
	if err := message.SendStartMessageToAdmin(s, m.GuildID, m.ChannelID, input); err != nil {
		message.SendErr(s, "起動通知をAdminサーバーに送信できません", m.GuildID, m.ChannelID, err)
	}

	// チャンネル一覧に追加
	shared.SetNewProcess(m.GuildID)
	defer shared.DeleteProcess(m.GuildID)

	// エントリーメッセージ
	msg, err := entry.SendEntryMessage(s, m, args.AnotherChannelID)
	if err != nil {
		message.SendErr(s, "エントリーメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	// カウントダウンメッセージ
	ok, err = countdown.CountDownScenario(s, msg, args.AnotherChannelID)
	if err != nil {
		message.SendErr(s, "カウントダウンメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}
	if !ok {
		return
	}

	// 開始メッセージ
	usrs, err := start.SendStartMessage(s, msg, args.AnotherChannelID)
	if err != nil {
		message.SendErr(s, "開始メッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	if battle.IsCanceledCheckAndSleep(10, m.GuildID) {
		return
	}

	// バトルメッセージ
	if err = battleMessage.NormalBattleMessageScenario(s, usrs, msg, args.AnotherChannelID); err != nil {
		message.SendErr(s, "バトルメッセージを送信できません", m.GuildID, m.ChannelID, err)
		return
	}

	// 正常終了のメッセージを送信
	if err = message.SendNormalFinishMessageToAdmin(s, m.GuildID); err != nil {
		message.SendErr(s, "終了通知をAdminサーバーに送信できません", m.GuildID, m.ChannelID, err)
		return
	}
}
