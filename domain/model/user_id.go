package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/techstart35/battle-bot/shared/errors"
)

// ユーザーIDです
type UserID struct {
	value string
}

// ユーザーIDを新規作成します
func NewUserID(value string) (UserID, error) {
	i := UserID{}
	i.value = value

	if err := i.Validate(); err != nil {
		return UserID{}, errors.NewError("ユーザーIDの検証に失敗しました", err)
	}

	return i, nil
}

// UserIDの値を文字列で取得します
func (i UserID) String() string {
	return i.value
}

// UserIDを比較します
func (i UserID) Equal(ii UserID) bool {
	return i.value == ii.value
}

// 全体を検証します
func (i UserID) Validate() error {
	if err := validator.New().Var(i.value, "required"); err != nil {
		return errors.NewError("検証に失敗しました", err)
	}

	return nil
}
