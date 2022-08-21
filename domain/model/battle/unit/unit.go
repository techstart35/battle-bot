package unit

import (
	"github.com/techstart35/battle-bot/domain/model/battle/unit/user"
	"github.com/techstart35/battle-bot/shared/errors"
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

// 検証します
func (u Unit) validate() error {
	// 生存者の重複を確認
	sm := map[string]string{}
	for _, v := range u.survivor {
		if _, ok := sm[v.ID().String()]; ok {
			return errors.NewError("生存者に重複した値が存在しています")
		}
		// 値は使用しないため、空を入れる
		sm[v.ID().String()] = ""
	}

	// 死者の重複を確認
	dm := map[string]string{}
	for _, v := range u.dead {
		if _, ok := dm[v.ID().String()]; ok {
			return errors.NewError("死者に重複した値が存在しています")
		}
		// 値は使用しないため、空を入れる
		sm[v.ID().String()] = ""
	}

	// 生存者と死者の重複を確認
	for sk, _ := range sm {
		if _, ok := dm[sk]; ok {
			return errors.NewError("生存者と死者で重複が発生しています")
		}
	}

	return nil
}
