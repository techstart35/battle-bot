package shared

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

// リアクションした人を取得します
//
// botのリアクションは除外します。
//
// botしかリアクションしない場合は、戻り値のスライスは空となります。
func GetReactedUsers(s *discordgo.Session, entryMsg *discordgo.Message) ([]*discordgo.User, error) {
	users := make([]*discordgo.User, 0)

	botName := os.Getenv("BOT_NAME")

	// 最大1000人まで参加可能（10 * 100）
	for i := 0; i < 10; i++ {
		var afterID string

		switch i {
		case 0:
			afterID = ""
		default:
			afterID = users[len(users)-1].ID
		}

		us, err := s.MessageReactions(entryMsg.ChannelID, entryMsg.ID, "⚔️", 100, "", afterID)
		if err != nil {
			return users, errors.New(fmt.Sprintf("リアクションをしたユーザーを取得できません: %v", err))
		}

		if len(us) == 0 || len(us) == 1 && us[0].Username == botName {
			break
		}

		// botは除外する
		for _, u := range us {
			if u.Username != botName {
				users = append(users, u)
			}
		}
	}

	return users, nil
}
