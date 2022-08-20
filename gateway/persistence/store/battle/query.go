package battle

import (
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/domain/model/battle"
	"github.com/techstart35/battle-bot/shared/errors"
)

// クエリです
type Query struct{}

// クエリを生成します
func NewQuery() (*Query, error) {
	return &Query{}, nil
}

// ギルドIDからバトルを取得します
//
// コールする場合は、NotFoundErrのエラーハンドリングをしてください。
func (q *Query) FindByGuildID(guildID model.GuildID) (*battle.Battle, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	if btl, ok := store.battle[guildID.String()]; ok {
		return btl, nil
	}

	return &battle.Battle{}, errors.NotFoundErr
}

// 全てのバトルを取得します
//
// コールする側で NotFoundErr の検証をします。
func (q *Query) FindAll() ([]*battle.Battle, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	res := make([]*battle.Battle, 0)

	if len(store.battle) == 0 {
		return res, errors.NotFoundErr
	}

	for _, v := range store.battle {
		res = append(res, v)
	}

	return res, nil
}

// 新規起動停止フラグを取得します
func (q *Query) IsStartRejected() bool {
	return isStartRejected
}
