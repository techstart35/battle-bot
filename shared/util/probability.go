package util

import (
	"math/rand"
	"time"
)

// 指定した確率でtrueが返ります
//
// 引数には1-10までの数字を入れます。
//
// 1を入れると10%,10を入れると100%の確率でtrueが返ります。
func CustomProbability(num, seed int) bool {
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// bをシャッフルする
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(b), func(i, j int) { b[i], b[j] = b[j], b[i] })

	// 再度シャッフルする
	rand.Seed(int64(seed))
	rand.Shuffle(len(b), func(i, j int) { b[i], b[j] = b[j], b[i] })

	gb := b[:num]

	for _, v := range gb {
		if v == 1 {
			return true
		}
	}

	return false
}
