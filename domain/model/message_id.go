package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/techstart35/battle-bot/shared/errors"
)

// メッセージIDです
type MessageID struct {
	value string
}

// メッセージIDを新規作成します
func NewMessageID(value string) (MessageID, error) {
	i := MessageID{}
	i.value = value

	if err := i.Validate(); err != nil {
		return MessageID{}, errors.NewError("メッセージIDの検証に失敗しました", err)
	}

	return i, nil
}

// メッセージIDの値を文字列で取得します
func (i MessageID) String() string {
	return i.value
}

// メッセージIDを比較します
func (i MessageID) Equal(ii MessageID) bool {
	return i.value == ii.value
}

// ステータスの値が設定されているか判別します
func (i MessageID) IsEmpty() bool {
	return i.value == ""
}

// 全体を検証します
func (i MessageID) Validate() error {
	if err := validator.New().Var(i.value, "required"); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}
