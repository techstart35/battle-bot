package template

import "github.com/techstart35/battle-bot/discord/shared"

// ソロギミックのテンプレートです
var SoloTemplates = []string{
	"💥｜**%s** は自爆した",
	"💥｜**%s** はバナナの皮で滑って気絶した",
	"💥｜**%s** は迷子で行方不明になった",
	"💥｜**%s** は腐った肉を食べて腹痛で戦闘不能",
	"💥｜**%s** はさびしくて孤独死",
	"💥｜**%s** はクッキー食べて口の水分無くなって死亡",
	"💥｜**%s** は豆腐の角に頭をぶつけて即死",
	"💥｜**%s** はタンスに小指をぶつけてショック死",
	"💥｜**%s** は鳥のフンが頭に落ちてやる気を失う。脱落。",
	"💥｜**%s** は寝坊で試合に間に合わなかった",
	"💥｜**%s** は盗んだバイクで走り出して事故った",
	"💥｜**%s** は夏の暑さで溶けてしまった",
	"💥｜**%s** は冬の寒さで凍ってしまった",
	"💥｜**%s** はモンハンしすぎて夜ふか死",
	"💥｜**%s** は嫁から鬼電、即帰宅。",
}

// バトルギミックのテンプレートです
var BattleTemplates = []string{
	"⚔️｜**%s** は **%s** をシンプルに殴って倒した",
	"⚔️｜**%s** は **%s** をブロッコリーで殴って倒した",
	"⚔️｜**%s** は **%s** を食べた",
	"⚔️｜**%s** は **%s** をタイキックで倒した。",
	"⚔️｜**%s** は三輪車で **%s** をぶっ飛ばした",
	"⚔️｜**%s** は **%s** に千年殺しをお見舞いした。アーメン。",
	"⚔️｜**%s** は **%s** を落とし穴に落とした",
	"⚔️｜**%s** は **%s** に即効性の毒を盛った。さようなら。",
	"⚔️｜**%s** は **%s** を全力の膝カックンで倒した",
	"⚔️｜**%s** はラリアットで **%s**を倒した",
}

// ソロギミックのテンプレートをランダムに取得します
func GetRandomSoloTmpl() string {
	return SoloTemplates[shared.RandInt(1, len(SoloTemplates)+1)-1]
}

// バトルギミックのテンプレートをランダムに取得します
func GetRandomBattleTmpl() string {
	return BattleTemplates[shared.RandInt(1, len(BattleTemplates)+1)-1]
}
