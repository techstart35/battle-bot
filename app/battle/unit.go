package battle

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/app/battle/template"
	"github.com/techstart35/battle-bot/domain/model"
	domainBattle "github.com/techstart35/battle-bot/domain/model/battle"
	"github.com/techstart35/battle-bot/domain/model/battle/unit"
	"github.com/techstart35/battle-bot/domain/model/battle/unit/user"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/util"
	"reflect"
	"strings"
	"time"
)

// ユニットのシナリオです
//
// コールする側で isCanceledErr のハンドリングを行います。
func (a *BattleApp) unitScenario(guildID model.GuildID) error {
	round := 1
	canRevive := false

	// 1回のループで、1つのUnitメッセージが送信されます
	//
	// 状態は毎回クエリで取得します
	for {
		// クエリー
		btl, err := a.Query.FindByGuildID(guildID)
		if err != nil {
			return errors.NewError("ギルドIDでバトルを取得できません", err)
		}

		svNum := len(btl.Unit().Survivor())
		stage := make([]user.User, 0)

		// 初回 & 生き残りが1名超の場合にsleepを入れます
		{
			if round == 1 && svNum > 1 {
				time.Sleep(10 * time.Second)
			}
		}

		// キャンセルを確認します
		if btl.IsCanceled() {
			return isCanceledErr
		}

		switch {
		// 生き残りが1名になった時点で、Winnerメッセージを送信
		case svNum == 1:
			time.Sleep(2 * time.Second)
			if err = a.sendWinnerMsgToUser(
				btl.Unit().Survivor()[0],
				btl.ChannelID(),
				btl.AnotherChannelID(),
			); err != nil {
				return errors.NewError("Winnerメッセージを送信できません", err)
			}
			return nil
		case svNum <= 12:
			// 全員をステージング
			stage = append(stage, btl.Unit().Survivor()...)
		case 12 < svNum && svNum < 60:
			// 12名をステージング
			for i, v := range btl.Unit().Survivor() {
				if i > 11 {
					break
				}
				stage = append(stage, v)
			}
		case svNum >= 60:
			// 20名をステージング
			for i, v := range btl.Unit().Survivor() {
				if i > 19 {
					break
				}
				stage = append(stage, v)
			}
			canRevive = false
		}

		// ユニットメッセージを作成
		res, err := a.createUnitMsg(stage)
		if err != nil {
			return errors.NewError("ユニットメッセージを作成できません", err)
		}

		// バトルを更新して永続化します
		//
		// Battle構造体も上書きします
		{
			btl, err = updateBattleByLoser(btl, res.Loser, round)
			if err != nil {
				return errors.NewError("バトルを更新できません", err)
			}

			if err = a.Repo.Update(btl); err != nil {
				return errors.NewError("更新できません", err)
			}
		}

		// ユニットメッセージを送信
		if err = a.sendUnitMsg(
			btl.ChannelID(),
			btl.AnotherChannelID(),
			res.Description,
			round,
			len(btl.Unit().Survivor()),
		); err != nil {
			return errors.NewError("ユニットメッセージを送信できません", err)
		}

		// 復活
		{
			// 死者が1名未満、または生き残りが2名以下の場合は復活イベントは発生しない
			if len(btl.Unit().Dead()) < 1 || len(btl.Unit().Survivor()) <= 2 {
				canRevive = false
			}

			var isRevived bool
			// 復活イベント
			if canRevive {
				revival, err := a.revivalScenario(
					btl.ChannelID(),
					btl.AnotherChannelID(),
					btl.Unit().Dead(),
				)
				if err != nil {
					return errors.NewError("復活イベントを起動できません", err)
				}

				// 復活イベントが送信された場合、集計して永続化します
				if !reflect.DeepEqual(revival, user.User{}) {
					b, err := updateBattleByRevive(btl, revival, round)
					if err != nil {
						return errors.NewError("バトルを更新できません", err)
					}

					if err = a.Repo.Update(b); err != nil {
						return errors.NewError("更新できません", err)
					}
					isRevived = true
				}
			}

			// 今回復活した場合は、次回の復活無し
			if isRevived {
				canRevive = false
			} else {
				canRevive = true
			}
		}

		if len(btl.Unit().Survivor()) > 1 {
			time.Sleep(17 * time.Second)
		}

		round++
	}
}

// ----------------------------------------
// Winnerメッセージの送信
// ----------------------------------------

// Winnerのメッセージを送信します
func (a *BattleApp) sendWinnerMsgToUser(
	winner user.User,
	chID model.ChannelID,
	anChID model.AnotherChannelID,
) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       "👑 Winner 👑",
		Description: fmt.Sprintf("<@%s>", winner.ID().String()),
		Color:       shared.ColorRed,
	}

	// エントリーチャンネルにメッセージを送信
	{
		_, err := a.Session.ChannelMessageSendEmbed(chID.String(), embedInfo)
		if err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}

		msg, err := a.Session.ChannelMessageSend(
			chID.String(),
			fmt.Sprintf("<@%s>さん、おめでとうございます🎉", winner.ID().String()),
		)
		if err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}

		if err = a.Session.MessageReactionAdd(
			msg.ChannelID, msg.ID, "🎉",
		); err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}
	}

	// 別チャンネルにメッセージを送信
	if !anChID.IsEmpty() {
		msg, err := a.Session.ChannelMessageSendEmbed(anChID.String(), embedInfo)
		if err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}

		if err = a.Session.MessageReactionAdd(
			msg.ChannelID, msg.ID, "🎉",
		); err != nil {
			return errors.NewError("リアクションを付与できません", err)
		}
	}

	return nil
}

// ----------------------------------------
// Unitを作成: winner,loserを集計
// ----------------------------------------

// battleを更新します
//
// 生き残りから敗者を除外し、死者に追加します。
func updateBattleByLoser(
	btl *domainBattle.Battle,
	loser []user.User,
	round int,
) (*domainBattle.Battle, error) {
	empty := &domainBattle.Battle{}

	survivor := make([]user.User, 0)
	// 既存の生き残りから敗者を削除
	{
		loserMap := map[string]user.User{}
		for _, v := range loser {
			loserMap[v.ID().String()] = v
		}

		for _, sv := range btl.Unit().Survivor() {
			if _, ok := loserMap[sv.ID().String()]; !ok {
				survivor = append(survivor, sv)
			}
		}
	}

	// 敗者を追加
	dead := append(btl.Unit().Dead(), loser...)

	r, err := unit.NewRound(uint(round) + 1)
	if err != nil {
		return empty, errors.NewError("ラウンドを作成できません", err)
	}

	u, err := unit.NewUnit(survivor, dead, r)
	if err != nil {
		return empty, errors.NewError("ユニットを作成できません", err)
	}

	btl.UpdateUnit(u)

	return btl, nil
}

// ----------------------------------------
// Unitメッセージの送信
// ----------------------------------------

// Unitメッセージを送信します
func (a *BattleApp) sendUnitMsg(
	chID model.ChannelID,
	anChID model.AnotherChannelID,
	message string,
	round, survivorNum int,
) error {
	// ユニットメッセージのテンプレートです
	const MessageTmpl = `
%s

生き残り: **%d名**
`

	embedInfo := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("第%d回戦", round),
		Description: fmt.Sprintf(MessageTmpl, message, survivorNum),
		Color:       shared.ColorOrange,
	}

	if !anChID.IsEmpty() {
		_, err := a.Session.ChannelMessageSendEmbed(anChID.String(), embedInfo)
		if err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}
	}

	_, err := a.Session.ChannelMessageSendEmbed(chID.String(), embedInfo)
	if err != nil {
		return errors.NewError("メッセージの送信に失敗しました", err)
	}
	return nil
}

// ----------------------------------------
// Unitメッセージの作成
// ----------------------------------------
const (
	battle     = "battle"
	soloBattle = "soloBattle"
	none       = "none"
)

// Unitメッセージ作成のレスポンスです
type CreateUnitMsgRes struct {
	Description string
	Loser       []user.User
}

// Unitのメッセージを作成します
func (a *BattleApp) createUnitMsg(stage []user.User) (CreateUnitMsgRes, error) {
	res := CreateUnitMsgRes{}

	// ユーザーのシャッフルを行います
	stg := make([]user.User, 0)
	for _, v := range stage {
		stg = append(stg, v)
	}
	stg = util.ShuffleUser(stg)

	loser := make([]user.User, 0)
	line := make([]string, 0)
	count := 0

	// 1回のループで、1文が作成されます。
	for {
		kind := a.getBattleKind(len(stg), count)
		switch kind {
		case battle:
			w := stg[0] // index[0]をwinnerとします
			l := stg[1] // index[1]をloserとします
			li := template.GetRandomBattleTmpl(
				w.Name().String(),
				l.Name().String(),
				count,
			)

			// 結果を追加
			{
				loser = append(loser, l)
				line = append(line, li)
			}

			// stgから2名を削除
			{
				for i := 0; i < 2; i++ {
					s, err := util.RemoveUserByIndex(stg, 0)
					if err != nil {
						return res, errors.NewError("Userのスライスから指定のindexを削除できません", err)
					}
					stg = s
				}
			}
		case soloBattle:
			l := stg[0]
			li := template.GetRandomSoloBattleTmpl(
				l.Name().String(),
				count,
			)

			// 結果を追加
			{
				loser = append(loser, l)
				line = append(line, li)
			}

			// stgから1名を削除
			{
				s, err := util.RemoveUserByIndex(stg, 0)
				if err != nil {
					return res, errors.NewError("Userのスライスから指定のindexを削除できません", err)
				}
				stg = s
			}
		case none:
			w := stg[0]
			li := template.GetRandomNoneTmpl(
				w.Name().String(),
				count,
			)

			// 結果を追加
			{
				line = append(line, li)
			}

			// stgから1名を削除
			{
				s, err := util.RemoveUserByIndex(stg, 0)
				if err != nil {
					return res, errors.NewError("Userのスライスから指定のindexを削除できません", err)
				}
				stg = s
			}
		default:
			return res, errors.NewError("バトルの種類が指定の値ではありません")
		}

		if len(stg) == 0 {
			break
		}

		// カウントアップ
		count++
	}

	res.Description = strings.Join(line, "\n")
	res.Loser = loser

	return res, nil
}

// バトルの種類を取得します
func (a *BattleApp) getBattleKind(stageNum, count int) string {
	// 最初の2名は必ずバトルとします
	if count == 0 {
		return battle
	}

	kind := soloBattle

	// 2人以上いる場合にkindの選択をします
	if stageNum > 1 {
		prob := map[string]int{
			battle:     40,
			soloBattle: 30,
			none:       30,
		}
		kind = util.ProbWithWeight(prob, count)
	}

	return kind
}

// battle,soloBattle,none作成時のレスポンスです
type FuncRes struct {
	winner []user.User
	loser  []user.User
	line   []string
	stage  []user.User
}

// ----------------------------------------
// 復活イベント
// ----------------------------------------

// 復活イベントのシナリオです
func (a *BattleApp) revivalScenario(
	chID model.ChannelID,
	anChID model.AnotherChannelID,
	dead []user.User,
) (user.User, error) {
	empty := user.User{}

	if len(dead) == 0 {
		return empty, nil
	}

	res := util.ProbWithWeight(map[string]int{
		"revival": 20,
		"none":    80,
	}, 0)

	if res == "revival" {
		time.Sleep(3 * time.Second)

		revival := util.ShuffleUser(dead)[0]

		if err := a.sendRevivalMessage(revival, chID, anChID); err != nil {
			return empty, errors.NewError("復活メッセージを送信できません")
		}

		return revival, nil
	}

	return empty, nil
}

// 復活メッセージを送信します
func (a *BattleApp) sendRevivalMessage(
	revival user.User,
	chID model.ChannelID,
	anChID model.AnotherChannelID,
) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       "敗者復活",
		Description: template.GetRandomRevivalTmpl(revival.Name().String()),
		Color:       shared.ColorPink,
	}

	_, err := a.Session.ChannelMessageSendEmbed(chID.String(), embedInfo)
	if err != nil {
		return errors.NewError("メッセージの送信に失敗しました", err)
	}

	if !anChID.IsEmpty() {
		_, err = a.Session.ChannelMessageSendEmbed(anChID.String(), embedInfo)
		if err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}
	}

	return nil
}

// 復活による生き残り、死者を集計してバトルを更新します
func updateBattleByRevive(
	btl *domainBattle.Battle,
	revival user.User,
	round int,
) (*domainBattle.Battle, error) {
	empty := &domainBattle.Battle{}

	survivor := append(btl.Unit().Survivor(), revival)
	dead, err := util.RemoveUserFromUsers(btl.Unit().Dead(), revival)
	if err != nil {
		return empty, errors.NewError("ユーザーのスライスからユーザーを削除できません", err)
	}

	r, err := unit.NewRound(uint(round) + 1)
	if err != nil {
		return empty, errors.NewError("ラウンドを作成できません", err)
	}

	u, err := unit.NewUnit(survivor, dead, r)
	if err != nil {
		return empty, errors.NewError("ユニットを作成できません", err)
	}

	btl.UpdateUnit(u)

	return btl, nil
}
