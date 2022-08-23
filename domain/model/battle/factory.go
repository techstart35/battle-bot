package battle

import (
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/shared/errors"
)

// Battle構造体を作成します
func BuildBattle(guildID, channelID, anotherChannelID, authorID string) (*Battle, error) {
	res := &Battle{}

	gID, err := model.NewGuildID(guildID)
	if err != nil {
		return res, errors.NewError("ギルドIDを作成できません", err)
	}

	cID, err := model.NewChannelID(channelID)
	if err != nil {
		return res, errors.NewError("チャンネルIDを作成できません", err)
	}

	anID, err := model.NewAnotherChannelID(anotherChannelID)
	if err != nil {
		return res, errors.NewError("配信チャンネルIDを作成できません", err)
	}

	author, err := model.NewUserID(authorID)
	if err != nil {
		return res, errors.NewError("起動者IDを作成できません", err)
	}

	res, err = NewBattle(gID, cID, anID, author)
	if err != nil {
		return res, errors.NewError("バトルを作成できません", err)
	}

	return res, nil
}
