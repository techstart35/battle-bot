package shared

import (
	"sync"
)

// チャンネル一覧です
//
// 一時停止指示が出ていた場合はTrueになります。
var process = sync.Map{}

func SetProcess(channelID string) {
	process.Store(channelID, true)
}

// プロセスの一覧を取得します
func GetProcess() []string {
	res := make([]string, 0)

	process.Range(func(key interface{}, value interface{}) bool {
		res = append(res, key.(string))
		return true
	})

	return res
}

// キャンセルされているかを確認します
func IsCanceled(channelID string) bool {
	if _, ok := process.Load(channelID); ok {
		return false
	}

	return true
}

// プロセスが起動中か確認します
func IsProcessing(channelID string) bool {
	if _, ok := process.Load(channelID); ok {
		return true
	}

	return false
}

// キャンセルします
func ProcessDelete(channelID string) {
	process.Delete(channelID)
}

// 新規起動を停止するフラグです
//
// true: 新規起動禁止
var IsStartRejected = false
