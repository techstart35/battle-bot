package template

import "github.com/techstart35/battle-bot/discord/shared"

// ソロギミックのテンプレートです
var SoloTemplates = []string{
	"💥｜**%s** は自爆した",
	"💥｜**%s** はバナナの皮で滑って気絶した",
}

// バトルギミックのテンプレートです
var BattleTemplates = []string{
	"⚔️｜**%s** は **%s** を倒した",
	"⚔️｜**%s** は **%s** を突き飛ばした",
}

// ソロギミックのテンプレートをランダムに取得します
func GetRandomSoloTmpl() string {
	return SoloTemplates[shared.RandInt(1, len(SoloTemplates))-1]
}

// バトルギミックのテンプレートをランダムに取得します
func GetRandomBattleTmpl() string {
	return BattleTemplates[shared.RandInt(1, len(BattleTemplates))-1]
}
