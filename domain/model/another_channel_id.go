package model

// 配信チャンネルのIDです
type AnotherChannelID struct {
	value string
}

// チャンネルIDを新規作成します
func NewAnotherChannelID(value string) (AnotherChannelID, error) {
	i := AnotherChannelID{}
	i.value = value

	return i, nil
}

// AnotherChannelIDの値を文字列で取得します
func (i AnotherChannelID) String() string {
	return i.value
}

// AnotherChannelIDを比較します
func (i AnotherChannelID) Equal(ii AnotherChannelID) bool {
	return i.value == ii.value
}

// ステータスの値が設定されているか判別します
func (i AnotherChannelID) IsEmpty() bool {
	return i.value == ""
}
