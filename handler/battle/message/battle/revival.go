package battle

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/handler/battle/message/battle/template"
	"github.com/techstart35/battle-bot/shared"
	"github.com/techstart35/battle-bot/shared/errors"
	"github.com/techstart35/battle-bot/shared/util"
	"time"
)

// 復活イベントを起動
func ExecRevivalEvent(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	anotherChannelID string,
	losers []*discordgo.User,
) (*discordgo.User, error) {
	if len(losers) == 0 {
		return nil, nil
	}

	// 20%の確率でイベントが発生
	// seedは敗者数を設定。変更可。
	if util.CustomProbability(2, len(losers)) {
		var revival *discordgo.User

		// 敗者の中から1名を選択
		losers = util.ShuffleDiscordUsers(losers)
		revival = losers[0]

		time.Sleep(3 * time.Second)
		// メッセージ送信
		if err := SendRevivalMessage(s, entryMessage, revival, anotherChannelID); err != nil {
			if errors.IsCanceledErr(err) {
				return nil, errors.CancelErr
			}

			return nil, errors.NewError("復活メッセージの送信に失敗しました", err)
		}

		return revival, nil
	}

	return nil, nil
}

// 復活メッセージを送信します
func SendRevivalMessage(
	s *discordgo.Session,
	entryMessage *discordgo.Message,
	user *discordgo.User,
	anotherChannelID string,
) error {
	embedInfo := &discordgo.MessageEmbed{
		Title:       "敗者復活",
		Description: template.GetRandomRevivalTmpl(user),
		Color:       shared.ColorPink,
	}

	_, err := s.ChannelMessageSendEmbed(entryMessage.ChannelID, embedInfo)
	if err != nil {
		return errors.NewError("メッセージの送信に失敗しました", err)
	}

	if anotherChannelID != "" {
		_, err := s.ChannelMessageSendEmbed(anotherChannelID, embedInfo)
		if err != nil {
			return errors.NewError("メッセージの送信に失敗しました", err)
		}
	}

	return nil
}
