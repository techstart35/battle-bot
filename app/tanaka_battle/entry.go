package tanaka_battle

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
)

// エントリーメッセージのテンプレートです
const entryTmpl = `
-------------------
***TANAKA ver***
-------------------

💘️️主催者：<@%s>
💘️参加：❤️️にリアクション
`

// 配信chがあった場合のエントリーメッセージのテンプレートです
const entryTmplWithAnCh = `
-------------------
***TANAKA ver***
-------------------

💘️主催者：<@%s>
💘️参加：❤️️にリアクション
💘配信：<#%s>
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
		Title:       "⚔️ Battle Royale ⚔️ ",
		Description: fmt.Sprintf(entryTmpl, btl.AuthorID().String()),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://pbs.twimg.com/profile_images/1562034684900835330/7uANsDm6_400x400.jpg",
		},
		Color:     shared.ColorPink,
		Timestamp: shared.GetNowTimeStamp(),
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
		shared.HeartBasic,
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
