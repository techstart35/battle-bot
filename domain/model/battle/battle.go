package battle

import (
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/domain/model/battle/unit"
	"github.com/techstart35/battle-bot/shared/errors"
)

// バトルです
type Battle struct {
	guildID          model.GuildID
	channelID        model.ChannelID
	anotherChannelID model.ChannelID
	entryMessageID   model.MessageID
	isCanceled       bool
	unit             unit.Unit
}

// バトルを作成します
func NewBattle(
	guildID model.GuildID,
	channelID, anotherChannelID model.ChannelID,
) (*Battle, error) {
	b := &Battle{}
	b.guildID = guildID
	b.channelID = channelID
	b.anotherChannelID = anotherChannelID
	b.entryMessageID = model.MessageID{}
	b.isCanceled = false
	b.unit = unit.Unit{}

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
// 以下 Getter
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
func (b *Battle) AnotherChannelID() model.ChannelID {
	return b.anotherChannelID
}

// エントリーメッセージIDを取得します
func (b *Battle) EntryMessageID() model.MessageID {
	return b.entryMessageID
}

// キャンセルフラグを取得します
func (b *Battle) IsCanceled() bool {
	return b.isCanceled
}

// ユニットを取得します
func (b *Battle) Unit() unit.Unit {
	return b.unit
}
