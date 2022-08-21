package util

import (
	"errors"
	"github.com/techstart35/battle-bot/domain/model/battle/unit/user"
	"math/rand"
	"time"
)

// Userのスライスから指定のindexを削除します
func RemoveUserByIndex(s []user.User, index int) ([]user.User, error) {
	if index >= len(s) {
		return nil, errors.New("indexが不正な値です")
	}

	res := make([]user.User, 0)

	for i, v := range s {
		if i != index {
			res = append(res, v)
		}
	}

	return res, nil
}

// Userのスライスから指定のUserを削除します
func RemoveUserFromUsers(users []user.User, u user.User) ([]user.User, error) {
	res := make([]user.User, 0)

	for _, uu := range users {
		if !uu.ID().Equal(u.ID()) {
			res = append(res, uu)
		}
	}

	return res, nil
}

// スライスの中身ををシャッフルします
func ShuffleUser(slice []user.User) []user.User {
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

// stringのスライスをシャッフルします
func ShuffleString(slice []string, seed int) []string {
	s := slice
	rand.Seed(int64(seed))
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}
