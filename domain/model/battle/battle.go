package battle

import (
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/domain/model/battle/unit"
	"github.com/techstart35/battle-bot/shared/errors"
	"time"
)

// バトルです
type Battle struct {
	guildID          model.GuildID
	channelID        model.ChannelID
	anotherChannelID model.AnotherChannelID
	entryMessageID   model.MessageID
	authorID         model.UserID
	isCanceled       bool
	unit             unit.Unit
	created          time.Time
}

// バトルを作成します
func NewBattle(
	guildID model.GuildID,
	channelID model.ChannelID,
	anotherChannelID model.AnotherChannelID,
	authorID model.UserID,
) (*Battle, error) {
	b := &Battle{}
	b.guildID = guildID
	b.channelID = channelID
	b.anotherChannelID = anotherChannelID
	b.entryMessageID = model.MessageID{}
	b.authorID = authorID
	b.isCanceled = false
	b.unit = unit.Unit{}
	b.created = time.Now()

	return b, nil
}

// エントリーメッセージIDを設定します
//
// 一度しか設定できません。
func (b *Battle) SetEntryMessage(msgID model.MessageID) error {
	if !b.entryMessageID.IsEmpty() {
		return errors.NewError("エントリーメッセージIDが設定されています")
	}

	b.entryMessageID = msgID

	return nil
}

// キャンセルフラグを設定します
//
// false -> true の一方向のみ変更可能です。
func (b *Battle) Cancel() {
	b.isCanceled = true
}

// ユニットを更新します
func (b *Battle) UpdateUnit(unit unit.Unit) {
	b.unit = unit
}

// ------------------
// getter
// ------------------

// ギルドIDを取得します
func (b *Battle) GuildID() model.GuildID {
	return b.guildID
}

// チャンネルIDを取得します
func (b *Battle) ChannelID() model.ChannelID {
	return b.channelID
}

// 配信チャンネルのIDを取得します
func (b *Battle) AnotherChannelID() model.AnotherChannelID {
	return b.anotherChannelID
}

// エントリーメッセージIDを取得します
func (b *Battle) EntryMessageID() model.MessageID {
	return b.entryMessageID
}

// 起動者のIDを取得します
func (b *Battle) AuthorID() model.UserID {
	return b.authorID
}

// キャンセルフラグを取得します
func (b *Battle) IsCanceled() bool {
	return b.isCanceled
}

// ユニットを取得します
func (b *Battle) Unit() unit.Unit {
	return b.unit
}
