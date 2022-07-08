package message

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/message/template"
	"github.com/techstart35/battle-bot/discord/shared"
	"math/rand"
	"strings"
	"time"
)

// バトルメッセージ全体のテンプレートです
var BattleMessageTemplate = `
%s

生き残り: **%d名**
`

// バトルメッセージを送信します
func BattleMessageHandler(
	s *discordgo.Session,
	users []*discordgo.User,
	entryMessage *discordgo.Message,
	anotherChannelID string,
) error {
	// エントリーが無い場合はNoEntryのメッセージを送信します
	if len(users) == 0 {
		if err := SendNoEntryMessage(s, entryMessage, anotherChannelID); err != nil {
			return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
		}

		return nil
	}

	survivor := users

	round := 1
	for {
		shuffleSurvivor(survivor)

		l := len(survivor)
		switch {
		// 生き残りが1名になった時点で、Winnerメッセージを送信
		case l == 1:
			if err := SendWinnerMessage(s, entryMessage, survivor[0], anotherChannelID); err != nil {
				return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
			}

			return nil
		case l <= 8: // 8以下の場合は、全員をステージングして対戦
			var stage []*discordgo.User
			stage = append(stage, survivor...)

			// バトルメッセージを作成
			res, err := createBattleMessage(stage)
			if err != nil {
				return errors.New(fmt.Sprintf("バトルメッセージの作成に失敗しました: %v", err))
			}

			// 生き残りを減らす
			survivor = res.Winners

			// バトルメッセージに生き残り数を追加
			description := fmt.Sprintf(
				BattleMessageTemplate,
				res.Description,
				len(survivor),
			)

			// メッセージ送信
			if err := SendBattleMessage(s, entryMessage, description, round, anotherChannelID); err != nil {
				return errors.New(fmt.Sprintf("Battleメッセージの送信に失敗しました: %v", err))
			}
			// カウントUP
			round++
		case l >= 8: // 8以上の場合は、8名のみをステージングして対戦
			var stage []*discordgo.User
			stage = survivor[0:8]

			res, err := createBattleMessage(stage)
			if err != nil {
				return errors.New(fmt.Sprintf("バトルメッセージの作成に失敗しました: %v", err))
			}

			// 生き残りを減らす
			var newSurvivor []*discordgo.User
			newSurvivor = append(newSurvivor, res.Winners...)
			newSurvivor = append(newSurvivor, survivor[8:]...)
			survivor = newSurvivor

			// バトルメッセージに生き残り数を追加
			description := fmt.Sprintf(
				BattleMessageTemplate,
				res.Description,
				len(survivor),
			)

			// メッセージ送信
			if err := SendBattleMessage(s, entryMessage, description, round, anotherChannelID); err != nil {
				return errors.New(fmt.Sprintf("Battleメッセージの送信に失敗しました: %v", err))
			}

			// カウントUP
			round++
		}

		if len(survivor) > 1 {
			time.Sleep(10 * time.Second)
		}
	}
}

// 生き残りのスライスの中身ををシャッフルします
func shuffleSurvivor(slice []*discordgo.User) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}

// Battleのメッセージを送信します
func SendBattleMessage(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	description string,
	round int,
	anotherChannelID string,
) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("第%d回戦", round),
		Description: description,
		Color:       0xff0000,
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
	Description string
	Winners     []*discordgo.User
}

// バトルメッセージを作成します
//
// usersが2未満の場合はエラーを返します。
//
// 生存数はこの関数を使う側で設定します。
//
// 1人以上のWinnerを返すため、最初の2名は必ずバトルとなります。
func createBattleMessage(users []*discordgo.User) (CreateBattleLinesRes, error) {
	var res CreateBattleLinesRes

	if len(users) < 2 {
		return res, errors.New("メッセージ作成に必要なユーザー数が不足しています")
	}

	var (
		lines   []string
		winners []*discordgo.User
		losers  []*discordgo.User
	)

	nextUsersIndex := 0

	for {
		num := 1

		// 2つ取得可能な場合のみ、ランダムで取得する
		if nextUsersIndex+1 != len(users) {
			num = shared.RandInt(1, 3)
		}

		// 必ずWinnerを設定するため、最初の2名は必ずバトルとする
		if nextUsersIndex == 0 {
			num = 2
		}

		switch num {
		case 1:
			loser := users[nextUsersIndex]
			line := fmt.Sprintf(
				template.GetRandomSoloTmpl(),
				loser.Username,
			)

			lines = append(lines, line)
			losers = append(losers, loser)

			nextUsersIndex++
		case 2:
			winner := users[nextUsersIndex]
			loser := users[nextUsersIndex+1]

			line := template.GetRandomBattleTmpl(winner.Username, loser.Username)

			lines = append(lines, line)
			winners = append(winners, winner)

			nextUsersIndex += 2
		default:
			return res, errors.New("取得したギミック数が不正です")
		}

		if nextUsersIndex == len(users) {
			break
		}
	}

	res.Description = strings.Join(lines, "\n")
	res.Winners = winners

	return res, nil
}
