package battle

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/handler/battle/message/battle/template"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/util"
	"strings"
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
func SendBattleMessage(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
	description string,
	round int,
	anotherChannelID string,
) error {
	// キャンセル指示を確認
	if shared.IsCanceled(m.GuildID) {
		return nil
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("第%d回戦", round),
		Description: description,
		Color:       shared.ColorOrange,
	}

	if anotherChannelID != "" {
		_, err := s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embedInfo)
	if err != nil {
		return errors.NewError("メッセージの送信に失敗しました", err)
	}

	return nil
}

// バトルメッセージを作成するレスポンスです
type CreateBattleLinesRes struct {
	Description string
	Winners     []*discordgo.User
	Losers      []*discordgo.User
}

// バトルメッセージを作成します
//
// usersが2未満の場合はエラーを返します。
//
// 生存数はこの関数を使う側で設定します。
//
// 1人以上のWinnerを返すため、最初の2名は必ずバトルとなります。
func CreateBattleMessage(
	m *discordgo.MessageCreate,
	stage []*discordgo.User,
) (CreateBattleLinesRes, error) {
	var res CreateBattleLinesRes

	// キャンセル指示を確認
	if shared.IsCanceled(m.GuildID) {
		return res, nil
	}

	if len(stage) < 2 {
		return res, errors.NewError("メッセージ作成に必要なユーザー数が不足しています")
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
		if shared.IsCanceled(m.GuildID) {
			return res, nil
		}

		num := soloNoBattle

		// 2つ取得可能な場合のみ、ランダムで取得する
		if nextUsersIndex+1 != len(stage) {
			tmpWaitList := []int{
				// soloBattle: 30%
				// soloNoBattle: 30%
				// battle: 40%
				soloBattle,
				soloNoBattle,
				battle,
				soloBattle,
				battle,
				soloNoBattle,
				battle,
				soloBattle,
				soloNoBattle,
				battle,
			}

			// ランダムにするため、スライスを2回シャッフル
			wl := util.ShuffleInt(tmpWaitList, nextUsersIndex)
			wl = util.ShuffleInt(wl, len(ls))

			num = wl[util.RandInt(1, 11)-1]
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

			line := template.GetRandomBattleTmpl(w.Username, l.Username, nextUsersIndex)

			lines = append(lines, line)
			ws = append(ws, w)
			ls = append(ls, l)

			nextUsersIndex += 2
		case soloNoBattle:
			w := stage[nextUsersIndex]
			line := fmt.Sprintf(
				template.GetRandomSoloTmpl(nextUsersIndex),
				w.Username,
			)

			lines = append(lines, line)
			// 負けていないため、勝者としてカウントする
			ws = append(ws, w)

			nextUsersIndex++
		default:
			return res, errors.NewError("取得したギミック数が不正です")
		}

		if nextUsersIndex == len(stage) {
			break
		}
	}

	res.Description = strings.Join(lines, "\n")
	res.Winners = ws
	res.Losers = ls

	return res, nil
}
