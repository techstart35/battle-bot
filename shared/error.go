package shared

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/message"
	"log"
)

// エラーをログ出力します
func LogErr(msg string, err error) {
	log.Println(fmt.Sprintf("%s: %v", msg, err))
}

// エラーを作成します
func CreateErr(msg string, err error) error {
	return fmt.Errorf("%s: %v", msg, err)
}

// エラーをTestチャンネルに送付します
func SendErr(s *discordgo.Session, msg, guildID, channelID string, err error) {
	var sendErrTmpl = `
【エラーが発生しました】
ギルドID: %s
チャンネルID: %s
メッセージ: %s
継承したエラー: %s
`
	m := fmt.Sprintf(sendErrTmpl, guildID, channelID, msg, err.Error())

	if err := message.SendSimpleEmbedMessage(s, "1003130506881277952", "", m); err != nil {
		panic("エラーメッセージを送信できません")
	}
}
