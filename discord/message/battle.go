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
		case l <= 16: // 16以下の場合は、全員をステージングして対戦
			var stage []*discordgo.User
			stage = append(stage, survivor...)

			// バトルメッセージを作成
			res, err := createBattleMessage(stage)
			if err != nil {
				return errors.New(fmt.Sprintf("バトルメッセージの作成に失敗しました: %v", err))
			}

			// メッセージ送信
			if err := SendBattleMessage(s, entryMessage, res.Description, round); err != nil {
				return errors.New(fmt.Sprintf("Battleメッセージの送信に失敗しました: %v", err))
			}

			// 生き残りを減らす
			survivor = res.Winners
			// カウントUP
			round++
		case l >= 16: // 16以上の場合は、16名のみをステージングして対戦
			var stage []*discordgo.User
			stage = survivor[0:16]

			res, err := createBattleMessage(stage)
			if err != nil {
				return errors.New(fmt.Sprintf("バトルメッセージの作成に失敗しました: %v", err))
			}

			// メッセージ送信
			if err := SendBattleMessage(s, entryMessage, res.Description, round); err != nil {
				return errors.New(fmt.Sprintf("Battleメッセージの送信に失敗しました: %v", err))
			}

			// 生き残りを減らす
			var newSurvivor []*discordgo.User
			newSurvivor = append(newSurvivor, res.Winners...)
			newSurvivor = append(newSurvivor, survivor[16:]...)

			survivor = newSurvivor
			// カウントUP
			round++
		}

		if len(survivor) > 1 {
			time.Sleep(5 * time.Second)
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
) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("第%d回戦", round),
		Description: description,
		Color:       0xff0000,
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
func createBattleMessage(users []*discordgo.User) (CreateBattleLinesRes, error) {
	var res CreateBattleLinesRes

	var (
		lines   []string
		winners []*discordgo.User
	)

	nextUsersIndex := 0

	for {
		num := 1

		// 2つ取得可能な場合のみ、ランダムで取得する
		if nextUsersIndex+1 != len(users) {
			num = shared.RandInt(1, 3)
		}

		switch num {
		case 1:
			line := fmt.Sprintf(
				template.GetRandomSoloTmpl(),
				shared.FormatMentionByUserID(users[nextUsersIndex].ID),
			)

			lines = append(lines, line)
			nextUsersIndex++
		case 2:
			line := fmt.Sprintf(
				template.GetRandomBattleTmpl(),
				shared.FormatMentionByUserID(users[nextUsersIndex].ID),
				shared.FormatMentionByUserID(users[nextUsersIndex+1].ID),
			)

			lines = append(lines, line)
			winners = append(winners, users[nextUsersIndex])

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
