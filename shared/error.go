package shared

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
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
	guildName, e := GetGuildName(s, guildID)
	if e != nil {
		log.Println("ギルドIDを取得できません", e)
	}

	var sendErrTmpl = `
ギルド名: %s

チャンネルID: %s

メッセージ: %s

継承したエラー: %s
`
	m := fmt.Sprintf(sendErrTmpl, guildName, channelID, msg, err.Error())

	if e := SendSimpleEmbedMessage(s, AdminChannelID, "エラーが発生しました", m, ColorRed); err != nil {
		LogErr("エラーメッセージをAdminサーバーに送信できません", e)
	}
}
