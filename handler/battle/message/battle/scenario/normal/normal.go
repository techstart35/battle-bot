package normal

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	battleMessage "github.com/techstart35/battle-bot/handler/battle/message/battle"
	"github.com/techstart35/battle-bot/handler/battle/message/noentry"
	"github.com/techstart35/battle-bot/handler/battle/message/winner"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/util"
	"time"
)

// バトルメッセージ全体のテンプレートです
var BattleMessageTemplate = `
%s

生き残り: **%d名**
`

const (
	BaseStageNum = 12
	NextStageNum = 20
)

// バトルメッセージを送信します
func NormalBattleMessageScenario(
	s *discordgo.Session,
	users []*discordgo.User,
	entryMessage *discordgo.Message,
	anotherChannelID string,
) error {
	// キャンセル指示を確認
	if shared.IsCanceled(entryMessage.GuildID) {
		return nil
	}

	var survivors []*discordgo.User
	var losers []*discordgo.User

	// エントリーが無い場合はNoEntryのメッセージを送信します
	if len(users) == 0 {
		if err := noentry.SendNoEntryMessage(s, entryMessage, anotherChannelID); err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}

		return nil
	}

	survivors = users

	round := 1
	for {
		// キャンセル指示を確認
		if shared.IsCanceled(entryMessage.GuildID) {
			return nil
		}

		survivors = util.ShuffleDiscordUsers(survivors)

		survivorLen := len(survivors)
		switch {
		// 生き残りが1名になった時点で、Winnerメッセージを送信
		case survivorLen == 1:
			time.Sleep(2 * time.Second)
			if err := winner.SendWinnerMessage(s, entryMessage, survivors[0], anotherChannelID); err != nil {
				return errors.NewError("メッセージの送信に失敗しました", err)
			}

			return nil

		// 基準数以下の場合は、全員をステージングして対戦
		case survivorLen <= BaseStageNum:
			var stage []*discordgo.User
			stage = append(stage, survivors...)

			// バトルメッセージを作成
			res, err := battleMessage.CreateBattleMessage(entryMessage, stage)
			if err != nil {
				return errors.NewError("バトルメッセージの送信に失敗しました", err)
			}

			// 生き残りと敗者を集計
			{
				// 生き残りを減らす
				survivors = res.Winners
				// 敗者を追加
				losers = append(losers, res.Losers...)
			}

			// バトルメッセージに生き残り数を追加
			description := fmt.Sprintf(BattleMessageTemplate, res.Description, len(survivors))

			// メッセージ送信
			if err := battleMessage.SendBattleMessage(s, entryMessage, description, round, anotherChannelID); err != nil {
				return errors.NewError("バトルメッセージの送信に失敗しました", err)
			}

			// カウントUP
			round++

			// 復活イベントを作成
			if len(survivors) > 2 && len(losers) >= 1 {
				revival, err := battleMessage.ExecRevivalEvent(s, entryMessage, anotherChannelID, losers)
				if err != nil {
					return errors.NewError("復活イベントの起動に失敗しました", err)
				}

				// 生き残りと敗者を集計
				if revival != nil {
					// 選択した1名をsurvivorに移行
					survivors = append(survivors, revival)
					// 選択した1名を敗者から削除
					ls, err := util.RemoveUserFromUsers(losers, 0)
					if err != nil {
						return errors.NewError("勝者の削除に失敗しました", err)
					}
					losers = ls
				}
			}

		// 基準数より多く、60未満の場合は、基準数のみをステージングして対戦
		case BaseStageNum < survivorLen && survivorLen < 60:
			var stage []*discordgo.User
			stage = survivors[0:BaseStageNum]

			res, err := battleMessage.CreateBattleMessage(entryMessage, stage)
			if err != nil {
				return errors.NewError("バトルメッセージの作成に失敗しました", err)
			}

			// 生き残りと敗者を集計
			{
				// 生き残りを減らす
				var newSurvivor []*discordgo.User
				newSurvivor = append(newSurvivor, res.Winners...)
				newSurvivor = append(newSurvivor, survivors[BaseStageNum:]...)
				survivors = newSurvivor

				// 敗者を追加
				losers = append(losers, res.Losers...)
			}

			// バトルメッセージに生き残り数を追加
			description := fmt.Sprintf(BattleMessageTemplate, res.Description, len(survivors))

			// メッセージ送信
			if err := battleMessage.SendBattleMessage(s, entryMessage, description, round, anotherChannelID); err != nil {
				return errors.NewError("バトルメッセージの送信に失敗しました", err)
			}

			// カウントUP
			round++

			// 復活イベントを作成
			if len(survivors) > 2 && len(losers) >= 1 {
				revival, err := battleMessage.ExecRevivalEvent(s, entryMessage, anotherChannelID, losers)
				if err != nil {
					return errors.NewError("復活イベントの起動に失敗しました", err)
				}

				// 生き残りと敗者を集計
				if revival != nil {
					// 選択した1名をsurvivorに移行
					survivors = append(survivors, revival)
					// 選択した1名を敗者から削除
					ls, err := util.RemoveUserFromUsers(losers, 0)
					if err != nil {
						return errors.NewError("敗者の削除に失敗しました", err)
					}
					losers = ls
				}
			}

		// 60以上の場合は、次の基準値をステージングして対戦
		case 60 <= survivorLen:
			var stage []*discordgo.User
			stage = survivors[0:NextStageNum]

			res, err := battleMessage.CreateBattleMessage(entryMessage, stage)
			if err != nil {
				return errors.NewError("バトルメッセージの作成に失敗しました", err)
			}

			// 生き残りと敗者を集計
			{
				// 生き残りを減らす
				var newSurvivor []*discordgo.User
				newSurvivor = append(newSurvivor, res.Winners...)
				newSurvivor = append(newSurvivor, survivors[NextStageNum:]...)
				survivors = newSurvivor

				// 敗者を追加
				losers = append(losers, res.Losers...)
			}

			// バトルメッセージに生き残り数を追加
			description := fmt.Sprintf(BattleMessageTemplate, res.Description, len(survivors))

			// メッセージ送信
			if err := battleMessage.SendBattleMessage(s, entryMessage, description, round, anotherChannelID); err != nil {
				return errors.NewError("バトルメッセージの送信に失敗しました", err)
			}

			// カウントUP
			round++
		}

		if len(survivors) > 1 {
			time.Sleep(17 * time.Second)
		}
	}
}
