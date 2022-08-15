package normal

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/message"
	"strings"
)

// 送信メッセージのテンプレートです
const checkInputTmpl = `
コマンドが間違っているか、チャンネルの権限が不足しています。
`

// コマンドの引数を確認します
//
// チャンネルIDを返します。
//
// inputが1つの場合は空の文字列を返します。
func CheckInput(s *discordgo.Session, channelID string, input []string) (string, error) {
	if len(input) > 2 {
		t := strings.TrimLeft(input[1], "<#")
		anotherChannelID := strings.TrimRight(t, ">")

		// チャンネルIDが正しいことを検証
		if _, err := s.Channel(anotherChannelID); err != nil {
			// エラーメッセージを送信
			if err = message.SendSimpleEmbedMessage(s, channelID, "ERROR", checkInputTmpl, 0); err != nil {
				return "", errors.NewError("CheckInputメッセージの送信に失敗しました", err)
			}

			return "", nil
		}

		return anotherChannelID, nil
	}

	return "", nil
}
