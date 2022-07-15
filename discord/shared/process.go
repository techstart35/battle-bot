package shared

// チャンネル一覧です
//
// 一時停止指示が出ていた場合はTrueになります。
var IsProcessing = map[string]bool{}

// 新規起動を停止するフラグです
//
// true: 新規起動禁止
var IsStartRejected bool
