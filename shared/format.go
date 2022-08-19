package shared

import "fmt"

// UserIDをメンションのフォーマットに変換します
func FormatMentionByUserID(userID string) string {
	return fmt.Sprintf("<@%s>", userID)
}

// ChannelIDをチャンネルリンク(メンション)のフォーマットに変換します
func FormatChannelIDToLink(channelID string) string {
	return fmt.Sprintf("<#%s>", channelID)
}

// チャンネルリンクを作成します
func CreateChannelURL(guildID, channelID string) string {
	base := "https://discord.com/channels/%s/%s"

	return fmt.Sprintf(base, guildID, channelID)
}
