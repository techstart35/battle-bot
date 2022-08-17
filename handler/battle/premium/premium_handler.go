package premium

//import (
//	"github.com/bwmarrin/discordgo"
//	"github.com/techstart35/battle-bot/handler/battle"
//	battleMessage "github.com/techstart35/battle-bot/handler/battle/message/battle/scenario"
//	"github.com/techstart35/battle-bot/handler/battle/message/countdown"
//	"github.com/techstart35/battle-bot/handler/battle/message/entry"
//	"github.com/techstart35/battle-bot/handler/battle/message/start"
//	"github.com/techstart35/battle-bot/shared"
//	"github.com/techstart35/battle-bot/shared/errors"
//	"github.com/techstart35/battle-bot/shared/message"
//	"strconv"
//	"strings"
//)
//
//// Battleを実行します
////
//// channelID(#判定) > userID(@判定) > winner(数字判定)
////
//// 1: pb
//// 2: pb <#channelID>
//// 3: pb <@userID>
//// 4: pb 3w
//// 5: pb <#channelID> <@userID> 3w
//func PremiumBattleHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
//	input := strings.Split(m.Content, " ")
//
//	cmd := input[0]
//	// コマンドが一致するか確認します
//	if cmd != shared.Command().Start {
//		return
//	}
//
//	ok, err := battle.CheckBeforeStartAndSendMessage(s, m.GuildID, m.ChannelID)
//	if err != nil {
//		message.SendErr(s, "起動前のチェックができません", m.GuildID, m.ChannelID, err)
//		return
//	}
//	if !ok {
//		return
//	}
//
//	args, err := CheckInput(s, input)
//	if err != nil {
//		message.SendErr(s, "コマンドが正しくありません", m.GuildID, m.ChannelID, err)
//		return
//	}
//
//	// Adminサーバーに起動メッセージを送信します
//	//
//	// Notice: ここでエラーが発生しても処理は継続させます
//	if err := message.SendStartMessageToAdmin(s, m.GuildID, m.ChannelID, input); err != nil {
//		message.SendErr(s, "起動通知をAdminサーバーに送信できません", m.GuildID, m.ChannelID, err)
//	}
//
//	// チャンネル一覧に追加
//	shared.SetNewProcess(m.GuildID)
//	defer shared.DeleteProcess(m.GuildID)
//
//	// エントリーメッセージ
//	_, err = entry.SendEntryMessage(s, m, args.AnotherChannelID)
//	if err != nil {
//		message.SendErr(s, "エントリーメッセージを送信できません", m.GuildID, m.ChannelID, err)
//		return
//	}
//
//	// カウントダウンメッセージ
//	err = countdown.CountDownScenario(s, m, args.AnotherChannelID)
//	if err != nil {
//		message.SendErr(s, "カウントダウンメッセージを送信できません", m.GuildID, m.ChannelID, err)
//		return
//	}
//	if !ok {
//		return
//	}
//
//	// 開始メッセージ
//	usrs, err := start.SendStartMessage(s, m, args.AnotherChannelID)
//	if err != nil {
//		message.SendErr(s, "開始メッセージを送信できません", m.GuildID, m.ChannelID, err)
//		return
//	}
//
//	if battle.IsCanceledCheckAndSleep(10, m.GuildID) {
//		return
//	}
//
//	// バトルメッセージ
//	if err = battleMessage.NormalBattleMessageScenario(s, usrs, m, args.AnotherChannelID); err != nil {
//		message.SendErr(s, "バトルメッセージを送信できません", m.GuildID, m.ChannelID, err)
//		return
//	}
//
//	// 正常終了のメッセージを送信
//	if err = message.SendNormalFinishMessageToAdmin(s, m.GuildID); err != nil {
//		message.SendErr(s, "終了通知をAdminサーバーに送信できません", m.GuildID, m.ChannelID, err)
//		return
//	}
//}
//
//// コマンドのレスポンスです
//type InputRes struct {
//	AnotherChannelID string
//	WinnerNum        uint
//	TargetUsers      []*discordgo.User
//}
//
//// コマンドの引数を確認します
//func CheckInput(s *discordgo.Session, input []string) (InputRes, error) {
//	res := InputRes{}
//
//	var anotherChannelID string
//	var winnerNum uint
//	targetUsers := make([]*discordgo.User, 0)
//
//	if len(input) < 2 {
//		return res, nil
//	}
//
//	for i, arg := range input {
//		if i == 0 {
//			continue
//		}
//
//		// 1.チャンネルID
//		// 2.ユーザーID
//		// 3.勝者数
//		if strings.Contains(arg, "<#") {
//			if anotherChannelID != "" {
//				return res, errors.NewError("チャンネルが複数設定されています")
//			}
//
//			t := strings.TrimLeft(arg, "<#")
//			anotherChannelID = strings.TrimRight(t, ">")
//
//			// チャンネルIDが正しいことを検証
//			if _, err := s.Channel(anotherChannelID); err != nil {
//				return res, errors.NewError("チャンネルの権限またはチャンネル名が不正です。", err)
//			}
//		} else if strings.Contains(arg, "<@") { // ユーザーID
//			if len(targetUsers) >= 3 {
//				return res, errors.NewError("ユーザー数が上限を超えています(上限3名)")
//			}
//
//			u := strings.TrimLeft(arg, "<@")
//			userID := strings.TrimRight(u, ">")
//
//			user, err := s.User(userID)
//			if err != nil {
//				return res, errors.NewError("ユーザーが不正な値です")
//			}
//
//			targetUsers = append(targetUsers, user)
//
//		} else if strings.Contains(arg, "w") { // 勝者数
//			wnStr := strings.TrimRight(arg, "w")
//			wnInt, err := strconv.Atoi(wnStr)
//			if err != nil {
//				return res, errors.NewError("文字列から数値に変換できません", err)
//			}
//
//			winnerNum = uint(wnInt)
//		} else {
//			return res, errors.NewError("コマンドに無効な値が含まれています")
//		}
//	}
//
//	res.AnotherChannelID = anotherChannelID
//	res.WinnerNum = winnerNum
//	res.TargetUsers = targetUsers
//
//	return res, nil
//}
