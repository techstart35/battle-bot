package battle

import (
	"github.com/techstart35/battle-bot/domain/model/battle"
	"sync"
)

// 状態保存構造体です
//
// ギルドID: battle構造体
type Store struct {
	battle map[string]battle.Battle
	mu     sync.Mutex
}

// 状態の保存領域です
//
// 今後DBに変更を予定しています。
var store = Store{}

// 新規起動を停止するフラグです
var isStartRejected bool
