package shared

import (
	"sync"
)

// 起動中のプロセスです
//
//
var process = sync.Map{}

// プロセスを新規追加します
func SetNewProcess(guildID string) {
	process.Store(guildID, true)
}

// プロセスをキャンセルします
func CancelProcess(guildID string) {
	if value, ok := process.Load(guildID); ok {
		// すでにキャンセルされている場合はfalseを返す
		if value == false {
			return
		}

		process.Store(guildID, false)
	}
}

// プロセスを削除します
func DeleteProcess(guildID string) {
	process.Delete(guildID)
}

// プロセスの一覧を取得します
func GetProcess() map[string]bool {
	res := map[string]bool{}

	process.Range(func(guildID interface{}, ok interface{}) bool {
		res[guildID.(string)] = ok.(bool)
		return true
	})

	return res
}

// キャンセルされているかを確認します
func IsCanceled(guildID string) bool {
	if value, ok := process.Load(guildID); ok {
		if value == true {
			return false
		} else {
			return true
		}
	}

	return false
}

// プロセスが起動中か確認します
func IsProcessing(guildID string) bool {
	if value, ok := process.Load(guildID); ok {
		return value.(bool)
	}

	return false
}

// プロセスが存在しているか確認します
func IsProcessExists(guildID string) bool {
	if _, ok := process.Load(guildID); ok {
		return true
	}

	return false
}

// 新規起動を停止するフラグです
//
// true: 新規起動禁止
var IsStartRejected = false
