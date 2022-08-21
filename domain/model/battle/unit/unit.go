package unit

import (
	"github.com/techstart35/battle-bot/domain/model/battle/unit/user"
)

// ユニットです
type Unit struct {
	survivor  []user.User
	dead      []user.User
	nextRound Round
}

// ユニットを作成します
func NewUnit(survivor, dead []user.User, nextRound Round) (Unit, error) {
	u := Unit{}
	u.survivor = survivor
	u.dead = dead
	u.nextRound = nextRound

	return u, nil
}

// 生き残りを取得します
func (u Unit) Survivor() []user.User {
	return u.survivor
}

// 死者を取得します
func (u Unit) Dead() []user.User {
	return u.dead
}

// ラウンド数を取得します
func (u Unit) Round() Round {
	return u.nextRound
}
