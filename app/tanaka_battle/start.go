package tanaka_battle

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/domain/model/battle/unit"
	"github.com/techstart35/battle-bot/domain/model/battle/unit/user"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"math"
	"os"
	"strings"
)

// 開始メッセージのシナリオです
//
// エントリーが0名の場合はNoEntryメッセージを送信します。
//
// コールする側で NoEntryErr, IsCanceledErr のハンドリングをします。
func (a *BattleApp) entryMsgScenario(guildID model.GuildID) error {
	// クエリー
	btl, err := a.Query.FindByGuildID(guildID)
	if err != nil {
		return errors.NewError("ギルドIDでバトルを取得できません", err)
	}

	// キャンセルを確認します
	if btl.IsCanceled() {
		return isCanceledErr
	}

	users, err := a.getReactedUsers(btl.ChannelID(), btl.EntryMessageID())
	if err != nil {
		return errors.NewError("リアクションしたユーザーを取得できません", err)
	}

	// usersの重複を排除します
	// 重複検証用のmapです
	idToUser := map[string]*discordgo.User{}
	{
		for _, u := range users {
			idToUser[u.ID] = u
		}
	}

	// usersを更新して永続化
	{
		survivor := make([]user.User, 0)
		for _, u := range idToUser {
			us, err := user.BuildUser(u.ID, u.Username)
			if err != nil {
				return errors.NewError("ユーザーを作成できません", err)
			}

			survivor = append(survivor, us)
		}

		r, err := unit.NewRound(1)
		if err != nil {
			return errors.NewError("ラウンドを作成できません", err)
		}

		un, err := unit.NewUnit(survivor, []user.User{}, r)
		if err != nil {
			return errors.NewError("ユニットを作成できません", err)
		}

		btl.UpdateUnit(un)

		if err = a.Repo.Update(btl); err != nil {
			return errors.NewError("更新できません", err)
		}
	}

	var challengers []string
	for _, v := range btl.Unit().Survivor() {
		challengers = append(challengers, v.Name().String())
	}

	// 参加者が0名の場合はNoEntryメッセージを送信
	if len(challengers) == 0 {
		if err = a.sendNoEntryMsgToUser(
			btl.ChannelID(),
			btl.AnotherChannelID(),
		); err != nil {
			return errors.NewError("NoEntryメッセージを送信できません", err)
		}

		return noEntryErr
	}

	// 開始メッセージを送信します
	if err = a.sendStartMsgToUser(
		btl.ChannelID(),
		btl.AnotherChannelID(),
		challengers,
	); err != nil {
		return errors.NewError("開始メッセージを送信できません", err)
	}

	return nil
}

// 開始メッセージのテンプレートです
//
// 配信chは必ずこれが送信されます。
const startTmpl = `
💘️挑戦者(%d名）：%s
💘️勝者：**1名**
💘️勝率：**%v％**
`

// 開始メッセージを送信します
func (a *BattleApp) sendStartMsgToUser(
	chID model.ChannelID,
	anChID model.AnotherChannelID,
	userNames []string,
) error {
	userNum := len(userNames)

	userStr := "100名を超えたため省略"
	if userNum < 100 {
		userStr = strings.Join(userNames, ", ")
	}

	var probability float64 = 0
	if userNum > 0 {
		p := 1 / float64(userNum) * 100
		probability = math.Round(p*10) / 10
	}

	// エントリーチャンネルに送信
	embedInfo := &discordgo.MessageEmbed{
		Title:       "⚔️ Battle Start ⚔️",
		Description: fmt.Sprintf(startTmpl, userNum, userStr, probability),
		Image: &discordgo.MessageEmbedImage{
			URL: "https://pbs.twimg.com/media/FWPF4oRaQAANksl?format=jpg&name=medium",
		},
		Color: shared.ColorPink,
	}

	_, err := a.Session.ChannelMessageSendEmbed(chID.String(), embedInfo)
	if err != nil {
		return errors.NewError("メッセージの送信に失敗しました", err)
	}

	// 配信チャンネルに送信
	if !anChID.IsEmpty() {
		embedInfo.Description = fmt.Sprintf(
			startTmpl, userNum, userStr, probability,
		)

		_, err = a.Session.ChannelMessageSendEmbed(anChID.String(), embedInfo)
		if err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}

		return nil
	}

	return nil
}

// NoEntryメッセージを送信します
func (a *BattleApp) sendNoEntryMsgToUser(
	chID model.ChannelID,
	anChID model.AnotherChannelID,
) error {
	const MsgTmpl = "エントリーがありませんでした"

	embedInfo := &discordgo.MessageEmbed{
		Title:       "No Entry",
		Description: MsgTmpl,
		Color:       shared.ColorPink,
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

// リアクションした人を取得します
//
// botのリアクションは除外します。
//
// botしかリアクションしない場合は、戻り値のスライスは空となります。
func (a *BattleApp) getReactedUsers(
	chID model.ChannelID,
	entryMsgID model.MessageID,
) ([]*discordgo.User, error) {
	users := make([]*discordgo.User, 0)

	botName := os.Getenv("BOT_NAME")

	// 最大1000人まで参加可能（10 * 100）
	for i := 0; i < 10; i++ {
		var afterID string

		switch i {
		case 0:
			afterID = ""
		default:
			afterID = users[len(users)-1].ID
		}

		us, err := a.Session.MessageReactions(chID.String(), entryMsgID.String(), shared.HeartBasic, 100, "", afterID)
		if err != nil {
			return users, errors.NewError("リアクションを取得できません", err)
		}

		if len(us) == 0 || len(us) == 1 && us[0].Username == botName {
			break
		}

		// botは除外する
		for _, u := range us {
			if u.Username != botName {
				users = append(users, u)
			}
		}
	}

	return users, nil
}
