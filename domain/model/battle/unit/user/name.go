package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/techstart35/battle-bot/shared/errors"
)

// ユーザー名です
type Name struct {
	value string
}

// ユーザー名を新規作成します
func NewName(v string) (Name, error) {
	n := Name{}
	n.value = v

	if err := n.validate(); err != nil {
		return n, errors.NewError("検証に失敗しました")
	}

	return n, nil
}

// 値を返します
func (n Name) String() string {
	return n.value
}

// 値を比較します
func (n Name) Equal(nn Name) bool {
	return n.value == nn.value
}

// 値が設定されているか判別します
func (n Name) IsEmpty() bool {
	return n.value == ""
}

// 検証します
func (n Name) validate() error {
	if err := validator.New().Var(n.value, "required"); err != nil {
		return errors.NewError("値の検証に失敗しました", err)
	}

	return nil
}
