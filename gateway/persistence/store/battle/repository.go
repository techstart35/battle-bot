package battle

import (
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/domain/model/battle"
	"github.com/techstart35/battle-bot/shared/errors"
)

// リポジトリです
type Repository struct{}

// リポジトリを作成します
func NewRepository() (*Repository, error) {
	return &Repository{}, nil
}

// バトルを作成します
func (r *Repository) Create(btl *battle.Battle) error {
	// すでに存在している場合は作成できません
	{
		_, err := r.FindByGuildID(btl.GuildID())
		if err != errors.NotFoundErr {
			return errors.NewError("ギルドIDでバトルを取得できません", err)
		}
	}

	store.mu.Lock()
	defer store.mu.Unlock()

	store.battle[btl.GuildID().String()] = btl

	return nil
}

// バトルを更新します
func (r *Repository) Update(btl *battle.Battle) error {
	// 存在していない場合は更新できません
	{
		_, err := r.FindByGuildID(btl.GuildID())
		if err == errors.NotFoundErr {
			return errors.NewError("バトルが存在していません", err)
		}
		if err != nil {
			return errors.NewError("ギルドIDでバトルを取得できません", err)
		}
	}

	store.mu.Lock()
	defer store.mu.Unlock()

	store.battle[btl.GuildID().String()] = btl

	return nil
}

// バトルを削除します
func (r *Repository) Delete(guildID model.GuildID) error {
	// 事前確認
	{
		// NotFoundErrは想定していないためエラーハンドリングは行いません
		_, err := r.FindByGuildID(guildID)
		if err != nil {
			return errors.NewError("ギルドIDでバトルを取得できません", err)
		}
	}

	delete(store.battle, guildID.String())

	return nil
}

// 新規起動を停止します
func (r *Repository) RejectStart() {
	isStartRejected = true
}

// ギルドIDからバトルを取得します
//
// コールする場合は、NotFoundErrのエラーハンドリングをしてください。
func (r *Repository) FindByGuildID(guildID model.GuildID) (*battle.Battle, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	btl, ok := store.battle[guildID.String()]
	if !ok {
		return &battle.Battle{}, errors.NotFoundErr
	}

	return btl, nil
}

// 全てのバトルを取得します
func (r *Repository) FindAll() ([]*battle.Battle, error) {
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
