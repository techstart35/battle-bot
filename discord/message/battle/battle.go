package battle

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/message/battle/template"
	"github.com/techstart35/battle-bot/discord/message/noentry"
	"github.com/techstart35/battle-bot/discord/message/winner"
	"github.com/techstart35/battle-bot/discord/shared"
	"strings"
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
func BattleMessageHandler(
	s *discordgo.Session,
	users []*discordgo.User,
	entryMessage *discordgo.Message,
	anotherChannelID string,
) error {
	// キャンセル指示を確認
	if !shared.IsProcessing[entryMessage.ChannelID] {
		return nil
	}

	var (
		survivor []*discordgo.User
		losers   []*discordgo.User
	)

	// エントリーが無い場合はNoEntryのメッセージを送信します
	if len(users) == 0 {
		if err := noentry.SendNoEntryMessage(s, entryMessage, anotherChannelID); err != nil {
			return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}

		return nil
	}

	survivor = users

	round := 1
	for {
		// キャンセル指示を確認
		if !shared.IsProcessing[entryMessage.ChannelID] {
			return nil
		}

		survivor = shared.ShuffleDiscordUsers(survivor)

		survivorLen := len(survivor)
		switch {
		// 生き残りが1名になった時点で、Winnerメッセージを送信
		case survivorLen == 1:
			time.Sleep(2 * time.Second)
			if err := winner.SendWinnerMessage(s, entryMessage, survivor[0], anotherChannelID); err != nil {
				return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
			}

			return nil

		// 基準数以下の場合は、全員をステージングして対戦
		case survivorLen <= BaseStageNum:
			var stage []*discordgo.User
			stage = append(stage, survivor...)

			// バトルメッセージを作成
			res, err := createBattleMessage(entryMessage, stage)
			if err != nil {
				return errors.New(fmt.Sprintf("バトルメッセージの作成に失敗しました: %v", err))
			}

			// 生き残りと敗者を集計
			{
				// 生き残りを減らす
				survivor = res.winners
				// 敗者を追加
				losers = append(losers, res.losers...)
			}

			// バトルメッセージに生き残り数を追加
			description := fmt.Sprintf(BattleMessageTemplate, res.description, len(survivor))

			// メッセージ送信
			if err := sendBattleMessage(s, entryMessage, description, round, anotherChannelID); err != nil {
				return errors.New(fmt.Sprintf("Battleメッセージの送信に失敗しました: %v", err))
			}

			// カウントUP
			round++

			// 復活イベントを作成
			if len(survivor) > 1 && len(losers) >= 1 {
				revival, err := execRevivalEvent(s, entryMessage, anotherChannelID, losers)
				if err != nil {
					return errors.New(fmt.Sprintf("復活イベントの起動に失敗しました: %v", err))
				}

				// 生き残りと敗者を集計
				if revival != nil {
					// 選択した1名をsurvivorに移行
					survivor = append(survivor, revival)
					// 選択した1名を敗者から削除
					ls, err := shared.RemoveUserFromUsers(losers, 0)
					if err != nil {
						return errors.New(fmt.Sprintf("敗者の削除に失敗しました: %v", err))
					}
					losers = ls
				}
			}

		// 基準数より多く、60未満の場合は、基準数のみをステージングして対戦
		case BaseStageNum < survivorLen && survivorLen < 60:
			var stage []*discordgo.User
			stage = survivor[0:BaseStageNum]

			res, err := createBattleMessage(entryMessage, stage)
			if err != nil {
				return errors.New(fmt.Sprintf("バトルメッセージの作成に失敗しました: %v", err))
			}

			// 生き残りと敗者を集計
			{
				// 生き残りを減らす
				var newSurvivor []*discordgo.User
				newSurvivor = append(newSurvivor, res.winners...)
				newSurvivor = append(newSurvivor, survivor[BaseStageNum:]...)
				survivor = newSurvivor

				// 敗者を追加
				losers = append(losers, res.losers...)
			}

			// バトルメッセージに生き残り数を追加
			description := fmt.Sprintf(BattleMessageTemplate, res.description, len(survivor))

			// メッセージ送信
			if err := sendBattleMessage(s, entryMessage, description, round, anotherChannelID); err != nil {
				return errors.New(fmt.Sprintf("Battleメッセージの送信に失敗しました: %v", err))
			}

			// カウントUP
			round++

			// 復活イベントを作成
			if len(survivor) > 1 && len(losers) >= 1 {
				revival, err := execRevivalEvent(s, entryMessage, anotherChannelID, losers)
				if err != nil {
					return errors.New(fmt.Sprintf("復活イベントの起動に失敗しました: %v", err))
				}

				// 生き残りと敗者を集計
				if revival != nil {
					// 選択した1名をsurvivorに移行
					survivor = append(survivor, revival)
					// 選択した1名を敗者から削除
					ls, err := shared.RemoveUserFromUsers(losers, 0)
					if err != nil {
						return errors.New(fmt.Sprintf("敗者の削除に失敗しました: %v", err))
					}
					losers = ls
				}
			}

		case 60 <= survivorLen: // 60以上の場合は、次の基準値をステージングして対戦
			var stage []*discordgo.User
			stage = survivor[0:NextStageNum]

			res, err := createBattleMessage(entryMessage, stage)
			if err != nil {
				return errors.New(fmt.Sprintf("バトルメッセージの作成に失敗しました: %v", err))
			}

			// 生き残りと敗者を集計
			{
				// 生き残りを減らす
				var newSurvivor []*discordgo.User
				newSurvivor = append(newSurvivor, res.winners...)
				newSurvivor = append(newSurvivor, survivor[NextStageNum:]...)
				survivor = newSurvivor

				// 敗者を追加
				losers = append(losers, res.losers...)
			}

			// バトルメッセージに生き残り数を追加
			description := fmt.Sprintf(BattleMessageTemplate, res.description, len(survivor))

			// メッセージ送信
			if err := sendBattleMessage(s, entryMessage, description, round, anotherChannelID); err != nil {
				return errors.New(fmt.Sprintf("Battleメッセージの送信に失敗しました: %v", err))
			}

			// カウントUP
			round++
		}

		if len(survivor) > 1 {
			time.Sleep(17 * time.Second)
		}
	}
}

// バトルメッセージを送信します
func sendBattleMessage(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	description string,
	round int,
	anotherChannelID string,
) error {
	// キャンセル指示を確認
	if !shared.IsProcessing[entryMessage.ChannelID] {
		return nil
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("第%d回戦", round),
		Description: description,
		Color:       0xffa500,
	}

	if anotherChannelID != "" {
		_, err := s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
	}

	return nil
}

// バトルメッセージを作成するレスポンスです
type CreateBattleLinesRes struct {
	description string
	winners     []*discordgo.User
	losers      []*discordgo.User
}

// バトルメッセージを作成します
//
// usersが2未満の場合はエラーを返します。
//
// 生存数はこの関数を使う側で設定します。
//
// 1人以上のWinnerを返すため、最初の2名は必ずバトルとなります。
func createBattleMessage(entryMessage *discordgo.Message, stage []*discordgo.User) (CreateBattleLinesRes, error) {
	var res CreateBattleLinesRes

	// キャンセル指示を確認
	if !shared.IsProcessing[entryMessage.ChannelID] {
		return res, nil
	}

	if len(stage) < 2 {
		return res, errors.New("メッセージ作成に必要なユーザー数が不足しています")
	}

	var (
		lines []string
		ws    []*discordgo.User
		ls    []*discordgo.User
	)

	nextUsersIndex := 0

	const (
		soloBattle   = iota + 1 // 1
		battle       = iota + 1 // 2
		soloNoBattle = iota + 1 // 3
	)

	for {
		// キャンセル指示を確認
		if !shared.IsProcessing[entryMessage.ChannelID] {
			return res, nil
		}

		num := soloNoBattle

		// 2つ取得可能な場合のみ、ランダムで取得する
		if nextUsersIndex+1 != len(stage) {
			tmpWaitList := []int{
				// soloBattle: 30%
				soloBattle,
				soloBattle,
				soloBattle,
				// soloNoBattle: 30%
				soloNoBattle,
				soloNoBattle,
				soloNoBattle,
				// battle: 40%
				battle,
				battle,
				battle,
				battle,
			}

			// ランダムにするため、スライスをシャッフル
			wl := shared.ShuffleInt(tmpWaitList, nextUsersIndex)

			num = wl[shared.RandInt(1, 11)-1]
		}

		// 必ずWinnerを設定するため、最初の2名は必ずバトルとする
		if nextUsersIndex == 0 {
			num = battle
		}

		switch num {
		case soloBattle:
			l := stage[nextUsersIndex]
			line := fmt.Sprintf(
				template.GetRandomSoloBattleTmpl(),
				l.Username,
			)

			lines = append(lines, line)
			ls = append(ls, l)

			nextUsersIndex++
		case battle:
			w := stage[nextUsersIndex]
			l := stage[nextUsersIndex+1]

			line := template.GetRandomBattleTmpl(w.Username, l.Username)

			lines = append(lines, line)
			ws = append(ws, w)
			ls = append(ls, l)

			nextUsersIndex += 2
		case soloNoBattle:
			w := stage[nextUsersIndex]
			line := fmt.Sprintf(
				template.GetRandomSoloTmpl(),
				w.Username,
			)

			lines = append(lines, line)
			// 負けていないため、勝者としてカウントする
			ws = append(ws, w)

			nextUsersIndex++
		default:
			return res, errors.New("取得したギミック数が不正です")
		}

		if nextUsersIndex == len(stage) {
			break
		}
	}

	res.description = strings.Join(lines, "\n")
	res.winners = ws
	res.losers = ls

	return res, nil
}

// 復活イベントを起動
func execRevivalEvent(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	anotherChannelID string,
	losers []*discordgo.User,
) (*discordgo.User, error) {
	// キャンセル指示を確認
	if !shared.IsProcessing[entryMessage.ChannelID] {
		return nil, nil
	}

	if len(losers) == 0 {
		return nil, nil
	}

	// 20%の確率でイベントが発生
	if shared.CustomProbability(2) {
		var revival *discordgo.User

		// 敗者の中から1名を選択
		losers = shared.ShuffleDiscordUsers(losers)
		revival = losers[0]

		time.Sleep(3 * time.Second)
		// メッセージ送信
		if err := SendRevivalMessage(s, entryMessage, revival, anotherChannelID); err != nil {
			return nil, errors.New(fmt.Sprintf("復活メッセージの送信に失敗しました: %v", err))
		}

		return revival, nil
	}

	return nil, nil
}
