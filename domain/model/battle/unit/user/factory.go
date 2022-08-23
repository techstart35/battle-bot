package user

import (
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/shared/errors"
)

// ユーザーを作成するfactoryです
func BuildUser(userID, userName string) (User, error) {
	user := User{}

	uID, err := model.NewUserID(userID)
	if err != nil {
		return user, errors.NewError("ユーザーIDを作成できません", err)
	}

	uName, err := NewName(userName)
	if err != nil {
		return user, errors.NewError("ユーザー名を作成できません", err)
	}

	user, err = NewUser(uID, uName)
	if err != nil {
		return user, errors.NewError("ユーザーを作成できません", err)
	}

	return user, nil
}
