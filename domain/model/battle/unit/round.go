package unit

import (
	"github.com/techstart35/battle-bot/shared/errors"
)

// ラウンドです
type Round struct {
	value uint
}

// ラウンドを新規作成します
func NewRound(value uint) (Round, error) {
	i := Round{}
	i.value = value

	if err := i.Validate(); err != nil {
		return Round{}, errors.NewError("ラウンドの検証に失敗しました", err)
	}

	return i, nil
}

// Roundの値を文字列で取得します
func (r Round) Uint() uint {
	return r.value
}

// Roundを比較します
func (r Round) Equal(ii Round) bool {
	return r.value == ii.value
}

// 全体を検証します
func (r Round) Validate() error {
	if r.value == 0 {
		return errors.NewError("検証に失敗しました")
	}

	return nil
}
