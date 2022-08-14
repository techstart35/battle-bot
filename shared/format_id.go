package shared

import "fmt"

// UserIDをメンションのフォーマットに変換します
func FormatMentionByUserID(userID string) string {
	return fmt.Sprintf("<@%s>", userID)
}

// ChannelIDをメンションのフォーマットに変換します
func FormatMentionByChannelID(channelID string) string {
	return fmt.Sprintf("<#%s>", channelID)
}
