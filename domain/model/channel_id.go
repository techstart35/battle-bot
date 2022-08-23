package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/techstart35/battle-bot/shared/errors"
)

// チャンネルChannelIDです
type ChannelID struct {
	value string
}

// チャンネルIDを新規作成します
func NewChannelID(value string) (ChannelID, error) {
	i := ChannelID{}
	i.value = value

	if err := i.Validate(); err != nil {
		return ChannelID{}, errors.NewError("チャンネルIDの検証に失敗しました", err)
	}

	return i, nil
}

// ChannelIDの値を文字列で取得します
func (i ChannelID) String() string {
	return i.value
}

// ChannelIDを比較します
func (i ChannelID) Equal(ii ChannelID) bool {
	return i.value == ii.value
}

// 全体を検証します
func (i ChannelID) Validate() error {
	if err := validator.New().Var(i.value, "required"); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}
