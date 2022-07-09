package template

import (
	"fmt"
	"github.com/techstart35/battle-bot/discord/shared"
)

// ソロバトルギミックのテンプレートをランダムに取得します
func GetRandomSoloBattleTmpl() string {
	var tmpl = []string{
		"💥｜**%s** は間違えて自爆ボタンを押してしまった。ﾄﾞｶｰﾝ。",
		"💥｜**%s** はバナナの皮で滑って気絶。",
		"💥｜**%s** は迷子で行方不明に。",
		"💥｜**%s** は腐った肉を食べて腹痛。戦闘不能。",
		"💥｜**%s** は豆の食べ過ぎで破裂。おまめぇ。",
		"💥｜**%s** はMIRAKO.の可愛さで失神。",
		"💥｜**%s** は豆腐の角に頭をぶつけて即死。",
		"💥｜**%s** はタンスに小指をぶつけてショック死。",
		"💥｜**%s** は鳥のフンが頭に落ちてやる気を失う。脱落。",
		"💥｜**%s** は寝坊で試合に間に合わなかった。",
		"💥｜**%s** は盗んだバイクで走り出したが事故。",
		"💥｜**%s** は夏の暑さで溶けてしまった。",
		"💥｜**%s** は冬の寒さで凍ってしまった。",
		"💥｜**%s** はモンハンしすぎて夜ふか死。",
		"💥｜**%s** は嫁から鬼電、即帰宅。",
		"💥｜**%s** はカツラが飛んで追いかけて退場。",
		"💥｜**%s** はMIRAKOにお仕置きされてうれ死",
		"💥｜**%s** は八門遁甲を開門するも、対戦相手がいなかった。",
		"💥｜**%s** は体に宿る九尾を抜かれて瀕死状態に。",
		"💥｜**%s** はジュレまみれになった。溺死。",
		"💥｜**%s** は白シャツにカレーを飛ばして戦意喪失。",
	}

	return tmpl[shared.RandInt(1, len(tmpl)+1)-1]
}

// バトルギミックのテンプレートをランダムに取得します
func GetRandomBattleTmpl(winner, loser string) string {
	var tmpl = []string{
		fmt.Sprintf("⚔️｜👑**%s** は念能力を取得。百式観音を発動し **%s** を天国へ葬った。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は **%s** をブロッコリーで殴った。", winner, loser),
		fmt.Sprintf("⚔️｜**%s** は 👑**%s** に食べられてしまった。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は **%s** にタイキックをかました。", winner, loser),
		fmt.Sprintf("⚔️｜**%s** は 👑**%s** に三輪車で轢かれた。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は **%s** に千年殺しをお見舞いした。アーメン。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は **%s** を落とし穴に落とした。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は **%s** に即効性の毒を盛った。さようなら。", winner, loser),
		fmt.Sprintf("⚔️｜**%s** は 👑**%s** に膝カックンされてやる気を失った。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** はラリアットで **%s** を倒した。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は豆鉄砲で **%s** を撃ち抜いた。", winner, loser),
		fmt.Sprintf("⚔️｜**%s** は 👑**%s** にきゅうりで殴られ死亡。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は写輪眼を開眼。 **%s** は幻術にかけられた。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は北斗百烈拳で **%s** を倒した。ほあたぁ!!", winner, loser),
		fmt.Sprintf("⚔️｜南斗水鳥拳伝承者の 👑**%s** は奥義 飛翔白麗で **%s** を倒した", winner, loser),
		fmt.Sprintf("⚔️｜**%s** は 👑**%s** に秘孔つかれてあべ死。", loser, winner),
		fmt.Sprintf("⚔️｜**%s** は 👑**%s** の筋肉バスターで気絶。", loser, winner),
		fmt.Sprintf("⚔️｜**%s** は 👑**%s** の投げたじゃがいもに当たって死亡。", loser, winner),
	}

	return tmpl[shared.RandInt(1, len(tmpl)+1)-1]
}

// ソロプレイ（無駄アクション）のテンプレートをランダムに取得します。
func GetRandomSoloTmpl() string {
	var tmpl = []string{
		"☀️｜天気が良かったので、 **%s** はお散歩に出かけた。",
		"☀️｜**%s** はのんきに釣りをしている。",
		"☀️｜**%s** は元気玉を作ろうと両手を上にあげている。",
		"☀️｜**%s** はキャンプを楽しんでいる。",
		"☀️｜**%s** はバナナを食べている。",
		"☀️｜**%s** は豆の収穫をしている。ﾀﾉｼｲ!!",
	}

	return tmpl[shared.RandInt(1, len(tmpl)+1)-1]
}
