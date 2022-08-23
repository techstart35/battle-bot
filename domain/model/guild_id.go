package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/techstart35/battle-bot/shared/errors"
)

// ギルドIDです
type GuildID struct {
	value string
}

// ギルドIDを新規作成します
func NewGuildID(value string) (GuildID, error) {
	i := GuildID{}
	i.value = value

	if err := i.Validate(); err != nil {
		return GuildID{}, errors.NewError("ギルドIDの検証に失敗しました", err)
	}

	return i, nil
}

// GuildIDの値を文字列で取得します
func (i GuildID) String() string {
	return i.value
}

// GuildIDを比較します
func (i GuildID) Equal(ii GuildID) bool {
	return i.value == ii.value
}

// 全体を検証します
func (i GuildID) Validate() error {
	if err := validator.New().Var(i.value, "required"); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}
