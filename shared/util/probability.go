package util

import (
	"math/rand"
	"time"
)

// 指定した確率で値を返します
//
// 確率は合計で100になっている必要があります。
// 100になっていない場合は空の文字列を返します。
//
// 返したい文字列: 確率
func ProbWithWeight(m map[string]int, seed int) string {
	// バリデーション
	{
		p := 0
		for _, v := range m {
			p += v
		}
		if p != 100 {
			return ""
		}
	}

	b := make([]string, 0)

	// 確率の数だけ配列に入れる
	for str, prob := range m {
		for i := 0; i < prob; i++ {
			b = append(b, str)
		}
	}

	// bをシャッフルする
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(b), func(i, j int) { b[i], b[j] = b[j], b[i] })

	// 再度bをシャッフルする
	rand.Seed(int64(seed))
	rand.Shuffle(len(b), func(i, j int) { b[i], b[j] = b[j], b[i] })

	return b[0]
}

// 指定した確率でtrueが返ります
//
// 引数には1-100までの数字を入れます。
func Prob(num, seed int) bool {
	b := make([]int, 0)
	for i := 1; i <= 100; i++ {
		b = append(b, i)
	}

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
