package template

import (
	"fmt"
	"github.com/techstart35/battle-bot/shared/util"
)

// ソロバトルギミックのテンプレートをランダムに取得します
func GetRandomSoloBattleTmpl(loser string, seed int) string {
	var tmpl = []string{
		fmt.Sprintf("💔｜**%s** は田中のタイキックで死亡。", loser),
		fmt.Sprintf("💔｜**%s** はジュレまみれで死亡。", loser),
	}

	// スライスをシャッフルする
	s := util.ShuffleString(tmpl, seed)

	return s[util.RandInt(1, len(tmpl)+1)-1]
}

// バトルギミックのテンプレートをランダムに取得します
func GetRandomBattleTmpl(winner, loser string, seed int) string {
	var tmpl = []string{
		fmt.Sprintf("💘｜**%s** は田中の靴下を ~~**%s**~~ に投げつけた。", winner, loser),
	}

	// スライスをシャッフルする
	s := util.ShuffleString(tmpl, seed)

	return s[util.RandInt(1, len(tmpl)+1)-1]
}

// noneのテンプレートをランダムに取得します。
func GetRandomNoneTmpl(winner string, seed int) string {
	var tmpl = []string{
		fmt.Sprintf("💞｜**%s** は田中と散歩中です。", winner),
	}

	// スライスをシャッフルする
	s := util.ShuffleString(tmpl, seed)

	return s[util.RandInt(1, len(tmpl)+1)-1]
}

// 復活のテンプレートをランダムに取得します
func GetRandomRevivalTmpl(revival string) string {
	var tmpl = []string{
		fmt.Sprintf("⚰️｜** %s ** は田中を倒すべく復活した。", revival),
	}

	return tmpl[util.RandInt(1, len(tmpl)+1)-1]
}
