package unit

import (
	"github.com/techstart35/battle-bot/domain/model/battle/unit/user"
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
