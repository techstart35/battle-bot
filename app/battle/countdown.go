package battle

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/domain/model/battle"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"time"
)

// カウントダウンのシナリオです
//
// キャンセル指示を確認します。
//
// コールする側で isCanceledErr のハンドリングを行います。
func (a *BattleApp) countDownScenario(guildID model.GuildID) error {
	// クエリー
	btl, err := a.Query.FindByGuildID(guildID)
	if err != nil {
		return errors.NewError("ギルドIDでバトルを取得できません", err)
	}

	// 60秒sleep
	time.Sleep(60 * time.Second)
	if btl.IsCanceled() {
		return isCanceledErr
	}

	// 60秒後（残り60秒）にメッセージを送信
	if err = a.sendCountDownMessage(btl, 60); err != nil {
		return errors.NewError("60秒前カウントダウンメッセージを送信できません", err)
	}

	// 30秒sleep
	time.Sleep(30 * time.Second)
	if btl.IsCanceled() {
		return isCanceledErr
	}

	// 残り30秒アナウンス
	if err = a.sendCountDownMessage(btl, 30); err != nil {
		return errors.NewError("30秒前カウントダウンメッセージを送信できません", err)
	}

	// 20秒sleep
	time.Sleep(20 * time.Second)
	if btl.IsCanceled() {
		return isCanceledErr
	}

	// 残り10秒アナウンス
	if err = a.sendCountDownMessage(btl, 10); err != nil {
		return errors.NewError("10秒前カウントダウンメッセージを送信できません", err)
	}

	// 10秒sleep
	time.Sleep(10 * time.Second)
	if btl.IsCanceled() {
		return isCanceledErr
	}

	return nil
}

// 基本的なカウントダウンテンプレートです
//
// 配信chは必ずこれが送信されます。
const countdownTmpl = `
開始まで **%d秒**

⚔️｜対戦
💥｜自滅
☀️｜敗者なし

[エントリーはこちら](%s)
`

// 別チャンネルありの場合のカウントダウンテンプレートです
//
// エントリーchのみ使用されます。
const countdownTmplToEntryChWithAnotherCh = `
開始まで **%d秒**

⚔️｜対戦
💥｜自滅
☀️｜敗者なし

[エントリーはこちら](%s)

<#%s> でも配信中
`

// カウントダウンメッセージを送信します
//
// 本メッセージ送信前にキャンセル指示を確認するため、
// この関数内ではキャンセル確認を行いません。
func (a *BattleApp) sendCountDownMessage(btl *battle.Battle, second int) error {
	const entryBaseURL = "https://discord.com/channels/%s/%s/%s"

	secondToColor := map[int]int{
		60: shared.ColorBlue,
		30: shared.ColorGreen,
		10: shared.ColorYellow,
	}

	// 秒数のバリデーション
	if _, ok := secondToColor[second]; !ok {
		return errors.NewError("秒数が指定の値ではありません")
	}

	entryURL := fmt.Sprintf(
		entryBaseURL,
		btl.GuildID().String(),
		btl.ChannelID().String(),
		btl.EntryMessageID().String(),
	)

	// 別チャンネルが無い場合を想定
	embedInfo := &discordgo.MessageEmbed{
		Title:       "⚔️ Battle Royale ⚔️",
		Description: fmt.Sprintf(countdownTmpl, second, entryURL),
		Color:       secondToColor[second],
	}

	// 別チャンネルがあった場合
	if !btl.AnotherChannelID().IsEmpty() {
		// エントリーチャンネルに送信
		{
			embedInfo.Description = fmt.Sprintf(
				countdownTmplToEntryChWithAnotherCh,
				second,
				entryURL,
				btl.AnotherChannelID().String(),
			)

			_, err := a.Session.ChannelMessageSendEmbed(btl.ChannelID().String(), embedInfo)
			if err != nil {
				return errors.NewError("メッセージの送信に失敗しました", err)
			}
		}

		// 別チャンネルに送信
		{
			embedInfo.Description = fmt.Sprintf(countdownTmpl, second, entryURL)

			_, err := a.Session.ChannelMessageSendEmbed(btl.AnotherChannelID().String(), embedInfo)
			if err != nil {
				return errors.NewError("メッセージの送信に失敗しました", err)
			}
		}

		return nil
	}

	_, err := a.Session.ChannelMessageSendEmbed(btl.ChannelID().String(), embedInfo)
	if err != nil {
		return errors.NewError("メッセージの送信に失敗しました", err)
	}

	return nil
}
