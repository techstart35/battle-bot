package user

import (
	"github.com/techstart35/battle-bot/domain/model"
)

// ユーザーです
type User struct {
	id   model.UserID
	name Name
}

// ユーザーを作成します
func NewUser(id model.UserID, name Name) (User, error) {
	u := User{}
	u.id = id
	u.name = name

	return u, nil
}

// idを取得します
func (u User) ID() model.UserID {
	return u.id
}

// ユーザー名を取得します
func (u User) Name() Name {
	return u.name
}
