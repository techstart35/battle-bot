package countdown

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/handler/battle"
	"github.com/techstart35/battle-bot/shared/errors"
)

// カウントダウンのシナリオです
//
// キャンセルされている場合はfalseを返します。
func CountDownScenario(s *discordgo.Session, m *discordgo.Message, anotherChannelID string) (bool, error) {
	if battle.IsCanceledCheckAndSleep(60, m.GuildID) {
		return false, nil
	}

	// 60秒後（残り60秒）にメッセージを送信
	if err := SendCountDownMessage(s, m, 60, anotherChannelID); err != nil {
		return false, errors.NewError("60秒前カウントダウンメッセージを送信できません", err)
	}

	if battle.IsCanceledCheckAndSleep(30, m.GuildID) {
		return false, nil
	}

	// 残り30秒アナウンス
	if err := SendCountDownMessage(s, m, 30, anotherChannelID); err != nil {
		return false, errors.NewError("30秒前カウントダウンメッセージを送信できません", err)
	}

	if battle.IsCanceledCheckAndSleep(20, m.GuildID) {
		return false, nil
	}

	// 残り10秒アナウンス
	if err := SendCountDownMessage(s, m, 10, anotherChannelID); err != nil {
		return false, errors.NewError("10秒前カウントダウンメッセージを送信できません", err)
	}

	if battle.IsCanceledCheckAndSleep(10, m.GuildID) {
		return false, nil
	}

	return true, nil
}
