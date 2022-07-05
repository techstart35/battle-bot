package discord

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
	"time"
)

// バトルメッセージを送信します
func BattleMessageHandler(
	s *discordgo.Session,
	users []*discordgo.User,
	entryMessage *discordgo.Message,
) error {
	// エントリーが無い場合はNoEntryのメッセージを送信します
	if len(users) == 0 {
		if err := sendNoEntryMessage(s, entryMessage); err != nil {
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
			if err := sendWinnerMessage(s, entryMessage, survivor[0]); err != nil {
				return errors.New(fmt.Sprintf("メッセージの送信に失敗しました: %v", err))
			}

			return nil
		case l <= 24 && l%2 == 0: // 24以下の偶数の場合は、全員をステージングして対戦
			var stage []*discordgo.User
			stage = append(stage, survivor...)

			var battleLines []string
			var winner []*discordgo.User

			// 2つずつ抽出する
			for i := 0; i < len(stage)-1; i += 2 {
				battleLine := fmt.Sprintf(getRandomBattleTmpl(), stage[i].Username, stage[i+1].Username)
				battleLines = append(battleLines, battleLine)

				// 勝者を追加
				winner = append(winner, stage[0])
			}

			// メッセージ送信
			description := strings.Join(battleLines, "\n")
			if err := sendBattleMessage(s, entryMessage, description, round); err != nil {
				return errors.New(fmt.Sprintf("Battleメッセージの送信に失敗しました: %v", err))
			}

			// 生き残りを減らす
			survivor = winner
			// カウントUP
			round++
		case l <= 24 && l%2 != 0: // 24以下の奇数の場合は、全員をステージングして、1名はソロ
			var stage []*discordgo.User
			stage = append(stage, survivor...)

			var battleLines []string
			var winner []*discordgo.User

			// 2つずつ抽出する
			for i := 0; i < len(stage); i += 2 {
				// 最後の1つ（奇数のため余る）はソロのギミックが適用される
				if i == len(stage)-1 {
					battleLine := fmt.Sprintf(getRandomSoloTmpl(), stage[i].Username)
					battleLines = append(battleLines, battleLine)

					break
				}

				battleLine := fmt.Sprintf(getRandomBattleTmpl(), stage[i].Username, stage[i+1].Username)
				battleLines = append(battleLines, battleLine)

				// 勝者を追加
				winner = append(winner, stage[0])
			}

			// メッセージ送信
			description := strings.Join(battleLines, "\n")
			if err := sendBattleMessage(s, entryMessage, description, round); err != nil {
				return errors.New(fmt.Sprintf("Battleメッセージの送信に失敗しました: %v", err))
			}

			// 生き残りを減らす
			survivor = winner
			// カウントUP
			round++
		case l >= 24: // 24以上の場合は、24名のみをステージングして対戦
			var stage []*discordgo.User
			stage = survivor[0:24]

			var battleLines []string
			var winner []*discordgo.User

			// 2つずつ抽出する
			for i := 0; i < len(stage)-1; i += 2 {
				battleLine := fmt.Sprintf(getRandomBattleTmpl(), stage[i].Username, stage[i+1].Username)
				battleLines = append(battleLines, battleLine)

				// 勝者を追加
				winner = append(winner, stage[0])
			}

			// メッセージ送信
			description := strings.Join(battleLines, "\n")
			if err := sendBattleMessage(s, entryMessage, description, round); err != nil {
				return errors.New(fmt.Sprintf("Battleメッセージの送信に失敗しました: %v", err))
			}

			// 生き残りを減らす
			survivor = winner
			// カウントUP
			round++
		}

		time.Sleep(5 * time.Second)
	}
}

// 生き残りのスライスの中身ををシャッフルします
func shuffleSurvivor(slice []*discordgo.User) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}

// ソロテンプレートをランダムに取得します
func getRandomSoloTmpl() string {
	var soloTemplates = []string{
		"💥｜**%s** は自爆した",
		"💥｜**%s** はバナナの皮で滑って気絶した",
	}

	return soloTemplates[RandInt(1, len(soloTemplates))-1]
}

// バトルテンプレートをランダムに取得します
func getRandomBattleTmpl() string {
	var battleTemplates = []string{
		"⚔️｜**%s** は **%s** を倒した",
		"⚔️｜**%s** は **%s** を突き飛ばした",
	}

	return battleTemplates[RandInt(1, len(battleTemplates))-1]
}
