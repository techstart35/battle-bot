package unit

import (
	"github.com/techstart35/battle-bot/domain/model/battle/unit/user"
	"github.com/techstart35/battle-bot/shared/errors"
)

// ユニットです
type Unit struct {
	survivor  []user.User
	loser     []user.User
	nextRound Round
}

// ユニットを作成します
func NewUnit(survivor, loser []user.User, nextRound Round) (Unit, error) {
	u := Unit{}
	u.survivor = survivor
	u.loser = loser
	u.nextRound = nextRound

	if err := u.validate(); err != nil {
		return Unit{}, errors.NewError("検証に失敗しました", err)
	}

	return u, nil
}

// 生き残りを取得します
func (u Unit) Survivor() []user.User {
	return u.survivor
}

// 敗者を取得します
func (u Unit) Loser() []user.User {
	return u.loser
}

// ラウンド数を取得します
func (u Unit) Round() Round {
	return u.nextRound
}

// 検証します
func (u Unit) validate() error {
	if len(u.survivor) < 1 {
		return errors.NewError("生き残りが0名です")
	}

	return nil
}
