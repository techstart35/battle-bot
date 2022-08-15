package premium

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared/errors"
	"strconv"
	"strings"
)

// コマンドのレスポンスです
type InputRes struct {
	AnotherChannelID string
	WinnerNum        uint
	TargetUsers      []*discordgo.User
}

// コマンドの引数を確認します
func CheckInput(s *discordgo.Session, input []string) (InputRes, error) {
	res := InputRes{}

	var anotherChannelID string
	var winnerNum uint
	targetUsers := make([]*discordgo.User, 0)

	if len(input) < 2 {
		return res, nil
	}

	for i, arg := range input {
		if i == 0 {
			continue
		}

		// 1.チャンネルID
		// 2.ユーザーID
		// 3.勝者数
		if strings.Contains(arg, "<#") {
			if anotherChannelID != "" {
				return res, errors.NewError("チャンネルが複数設定されています")
			}

			t := strings.TrimLeft(arg, "<#")
			anotherChannelID = strings.TrimRight(t, ">")

			// チャンネルIDが正しいことを検証
			if _, err := s.Channel(anotherChannelID); err != nil {
				return res, errors.NewError("チャンネルの権限またはチャンネル名が不正です。", err)
			}
		} else if strings.Contains(arg, "<@") { // ユーザーID
			if len(targetUsers) >= 3 {
				return res, errors.NewError("ユーザー数が上限を超えています(上限3名)")
			}

			u := strings.TrimLeft(arg, "<@")
			userID := strings.TrimRight(u, ">")

			user, err := s.User(userID)
			if err != nil {
				return res, errors.NewError("ユーザーが不正な値です")
			}

			targetUsers = append(targetUsers, user)

		} else if strings.Contains(arg, "w") { // 勝者数
			wnStr := strings.TrimRight(arg, "w")
			wnInt, err := strconv.Atoi(wnStr)
			if err != nil {
				return res, errors.NewError("文字列から数値に変換できません", err)
			}

			winnerNum = uint(wnInt)
		} else {
			return res, errors.NewError("コマンドに無効な値が含まれています")
		}
	}

	res.AnotherChannelID = anotherChannelID
	res.WinnerNum = winnerNum
	res.TargetUsers = targetUsers

	return res, nil
}
