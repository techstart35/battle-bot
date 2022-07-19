package shared

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
)

// Userのスライスから指定のindexを削除します
func RemoveUserFromUsers(s []*discordgo.User, i int) ([]*discordgo.User, error) {
	if i >= len(s) {
		return nil, errors.New("indexが不正な値です")
	}

	var res []*discordgo.User
	res = s

	res = append(res[:i], res[i+1:]...)

	return res, nil
}

// スライスの中身ををシャッフルします
func ShuffleDiscordUsers(slice []*discordgo.User) []*discordgo.User {
	s := slice
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}

// intのスライスをシャッフルします
func ShuffleInt(slice []int, seed int) []int {
	s := slice
	rand.Seed(int64(seed))
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}
