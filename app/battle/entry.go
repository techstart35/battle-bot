package battle

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
)

// エントリーメッセージのテンプレートです
const entryTmpl = `
⚡️主催者：<@%s>
⚡️勝者：**1名**
⚡️開始：**2分後**
⚡️参加：⚔️にリアクション
`

// 配信chがあった場合のエントリーメッセージのテンプレートです
const entryTmplWithAnCh = `
⚡️主催者：<@%s>
⚡️勝者：**1名**
⚡️開始：**2分後**
⚡️参加：⚔️にリアクション
⚡️配信：<#%s>
`

// エントリーメッセージを送信します
//
// 起動元のチャンネルのみに送信します。
//
// この関数ではキャンセル処理の確認を行いません。
func (a *BattleApp) sendEntryMsgToUser(guildID model.GuildID) error {
	// クエリー
	btl, err := a.Query.FindByGuildID(guildID)
	if err != nil {
		return errors.NewError("ギルドIDでバトルを取得できません", err)
	}

	// キャンセルを確認します
	if btl.IsCanceled() {
		return isCanceledErr
	}

	embedInfo := &discordgo.MessageEmbed{
		Title:       "⚔️ Battle Royale ⚔️",
		Description: fmt.Sprintf(entryTmpl, btl.AuthorID().String()),
		Color:       shared.ColorCyan,
		Timestamp:   shared.GetNowTimeStamp(),
	}

	// 別チャンネルの指定があった場合はテンプレートを差し替え
	if !btl.AnotherChannelID().IsEmpty() {
		embedInfo.Description = fmt.Sprintf(
			entryTmplWithAnCh,
			btl.AuthorID().String(),
			btl.AnotherChannelID().String(),
		)
	}

	msg, err := a.Session.ChannelMessageSendEmbed(btl.ChannelID().String(), embedInfo)
	if err != nil {
		return errors.NewError("メッセージの送信に失敗しました", err)
	}

	// リアクションを付与
	if err = a.Session.MessageReactionAdd(
		btl.ChannelID().String(),
		msg.ID,
		"⚔️",
	); err != nil {
		return errors.NewError("リアクションを付与できません", err)
	}

	// エントリーメッセージIDを永続化します
	{
		entryMsgID, err := model.NewMessageID(msg.ID)
		if err != nil {
			return errors.NewError("エントリーメッセージIDを作成できません", err)
		}

		if err = btl.SetEntryMessage(entryMsgID); err != nil {
			return errors.NewError("エントリーメッセージIDを設定できません", err)
		}

		if err = a.Repo.Update(btl); err != nil {
			return errors.NewError("バトルを更新できません", err)
		}
	}

	return nil
}
