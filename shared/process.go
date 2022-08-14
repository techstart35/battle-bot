package shared

import (
	"sync"
)

// 起動中のプロセスです
//
//
var process = sync.Map{}

// プロセスを新規追加します
func SetNewProcess(channelID string) {
	process.Store(channelID, true)
}

// プロセスをキャンセルします
func CancelProcess(channelID string) {
	if value, ok := process.Load(channelID); ok {
		// すでにキャンセルされている場合はfalseを返す
		if value == false {
			return
		}

		process.Store(channelID, false)
	}
}

// プロセスを削除します
func DeleteProcess(channelID string) {
	process.Delete(channelID)
}

// プロセスの一覧を取得します
func GetProcess() map[string]bool {
	res := map[string]bool{}

	process.Range(func(key interface{}, value interface{}) bool {
		res[key.(string)] = value.(bool)
		return true
	})

	return res
}

// キャンセルされているかを確認します
func IsCanceled(channelID string) bool {
	if value, ok := process.Load(channelID); ok {
		if value == true {
			return false
		} else {
			return true
		}
	}

	return false
}

// プロセスが起動中か確認します
func IsProcessing(channelID string) bool {
	if _, ok := process.Load(channelID); ok {
		return true
	}

	return false
}

// 新規起動を停止するフラグです
//
// true: 新規起動禁止
var IsStartRejected = false
