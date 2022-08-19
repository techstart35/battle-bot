package countdown

//import (
//	"fmt"
//	"github.com/bwmarrin/discordgo"
//	"github.com/techstart35/battle-bot/handler/battle"
//	"github.com/techstart35/battle-bot/shared"
//	"github.com/techstart35/battle-bot/shared/errors"
//)
//
//// カウントダウンのシナリオです
////
//// キャンセル指示を確認します。
//func CountDownScenario(
//	s *discordgo.Session,
//	entryMessage *discordgo.Message,
//	guildID, anotherChannelID string,
//) error {
//	// 60秒sleep
//	if battle.IsCanceledCheckAndSleep(60, guildID) {
//		return errors.CancelErr
//	}
//
//	// 60秒後（残り60秒）にメッセージを送信
//	if err := SendCountDownMessage(s, entryMessage, 60, guildID, anotherChannelID); err != nil {
//		return errors.NewError("60秒前カウントダウンメッセージを送信できません", err)
//	}
//
//	// 30秒sleep
//	if battle.IsCanceledCheckAndSleep(30, guildID) {
//		return errors.CancelErr
//	}
//
//	// 残り30秒アナウンス
//	if err := SendCountDownMessage(s, entryMessage, 30, guildID, anotherChannelID); err != nil {
//		return errors.NewError("30秒前カウントダウンメッセージを送信できません", err)
//	}
//
//	// 20秒sleep
//	if battle.IsCanceledCheckAndSleep(20, guildID) {
//		return errors.CancelErr
//	}
//
//	// 残り10秒アナウンス
//	if err := SendCountDownMessage(s, entryMessage, 10, guildID, anotherChannelID); err != nil {
//		return errors.NewError("10秒前カウントダウンメッセージを送信できません", err)
//	}
//
//	// 10秒sleep
//	if battle.IsCanceledCheckAndSleep(10, guildID) {
//		return errors.CancelErr
//	}
//
//	return nil
//}
//
//// エントリーチャンネルに送信するカウントダウンメッセージです
//var entryChannelTemplate = `
//開始まで **%d秒**
//
//⚔️-対戦
//💥-自滅
//☀️-敗者なし
//
//[エントリーはこちら](%s)
//
//<#%s> でも配信中 💬
//`
//
//// エントリーチャンネルに送信するカウントダウンメッセージです
////
//// 別チャンネルが指定されていない場合に使用します。
//var noAnotherChannelTemplate = `
//開始まで **%d秒**
//
//⚔️-対戦
//💥-自滅
//☀️-敗者なし
//
//[エントリーはこちら](%s)
//`
//
//// 別チャンネルに送信するカウントダウンメッセージです
//var anotherChannelTemplate = `
//開始まで **%d秒**
//
//⚔️-対戦
//💥-自滅
//☀️-敗者なし
//
//[エントリーはこちら](%s)
//`
//
//const entryBaseURL = "https://discord.com/channels/%s/%s/%s"
//
//// カウントダウンメッセージを送信します
////
//// 本メッセージ送信前にキャンセル指示を確認するため、
//// この関数内ではキャンセル確認を行いません。
//func SendCountDownMessage(
//	s *discordgo.Session,
//	entryMessage *discordgo.Message,
//	beforeStart uint,
//	guildID, anotherChannelID string,
//) error {
//	var color int
//	switch beforeStart {
//	case 60:
//		color = shared.ColorBlue
//	case 30:
//		color = shared.ColorGreen
//	case 10:
//		color = shared.ColorYellow
//	}
//
//	entryURL := fmt.Sprintf(entryBaseURL, guildID, entryMessage.ChannelID, entryMessage.ID)
//
//	// 別チャンネルが無い場合を想定
//	embedInfo := &discordgo.MessageEmbed{
//		Title: "⚔️ Battle Royale ⚔️",
//		Description: fmt.Sprintf(
//			noAnotherChannelTemplate,
//			beforeStart,
//			entryURL,
//		),
//		Color: color,
//	}
//
//	// 別チャンネルがあった場合
//	if anotherChannelID != "" {
//		// エントリーチャンネルに送信
//		embedInfo.Description = fmt.Sprintf(
//			entryChannelTemplate,
//			beforeStart,
//			entryURL,
//			anotherChannelID,
//		)
//
//		_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
//		if err != nil {
//			return errors.NewError("メッセージの送信に失敗しました", err)
//		}
//
//		// 別チャンネルに送信
//		embedInfo.Description = fmt.Sprintf(
//			anotherChannelTemplate,
//			beforeStart,
//			entryURL,
//		)
//
//		_, err = s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
//		if err != nil {
//			return errors.NewError("メッセージの送信に失敗しました", err)
//		}
//
//		return nil
//	}
//
//	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
//	if err != nil {
//		return errors.NewError("メッセージの送信に失敗しました", err)
//	}
//
//	return nil
//}
