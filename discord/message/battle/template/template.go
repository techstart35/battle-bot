package template

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/discord/shared"
)

// ソロバトルギミックのテンプレートをランダムに取得します
func GetRandomSoloBattleTmpl() string {
	var tmpl = []string{
		"💥｜**%s** は間違えて自爆ボタンを押してしまった💥",
		"💥｜**%s** はバナナの皮で滑って気絶🍌",
		"💥｜**%s** は迷子で行方不明に。",
		"💥｜**%s** は腐った肉を食べて腹痛。戦闘不能。",
		"💥｜**%s** は豆の食べ過ぎで破裂。おまめぇ。",
		"💥｜**%s** は豆腐の角に頭をぶつけて即死。",
		"💥｜**%s** はタンスに小指をぶつけてショック死。",
		"💥｜**%s** は鳥のフンが頭に落ちてやる気を失う。脱落。",
		"💥｜**%s** は寝坊で試合に間に合わなかった。",
		"💥｜**%s** は盗んだバイクで走り出したが事故。",
		"💥｜**%s** は夏の暑さで溶けてしまった🫠",
		"💥｜**%s** は冬の寒さで凍ってしまった🥶",
		"💥｜**%s** は嫁から鬼電、即帰宅💨",
		"💥｜**%s** は飛んでいったカツラを追いかけて退場。",
		"💥｜**%s** は白シャツにカレーを飛ばして戦意喪失。",
		"💥｜**%s** は賞味期限切れの生卵を食す。腹を壊した。",
		"💥｜**%s** は快速特急に乗ってしまい下車するはずの駅で降りられず。脱落。",
		"💥｜**%s** はランブルの勝ち方を解明すべくアマゾンの奥地へ向かった。",
		"💥｜**%s** は木登りをしていたが、足を滑らせ滑落。",
		"💥｜**%s** はつま先立ちで歩いていたため、足の指を骨折。",
	}

	return tmpl[shared.RandInt(1, len(tmpl)+1)-1]
}

// バトルギミックのテンプレートをランダムに取得します
func GetRandomBattleTmpl(winner, loser string) string {
	var tmpl = []string{
		fmt.Sprintf("⚔️｜👑**%s** は念能力を取得。百式観音を発動し 💀**%s** を駆逐した。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** をブロッコリーで殴った🥦", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** に食べられてしまった。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** をタイキックで蹴り倒した。", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** に三輪車で轢かれた。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** に千年殺しをお見舞いした。アーメン。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** を落とし穴に落とした。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** に即効性の毒を盛った。さようなら。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** はラリアットで 💀**%s** を倒した。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は豆鉄砲で 💀**%s** を撃ち抜いた。", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** にきゅうりで殴られ死亡。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は写輪眼を開眼。 💀**%s** は幻術にかけられた。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は北斗百烈拳で 💀**%s** を倒した。ほあたぁ!!", winner, loser),
		fmt.Sprintf("⚔️｜南斗水鳥拳伝承者の 👑**%s** は奥義 飛翔白麗で 💀**%s** を倒した。", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** に秘孔つかれてあべ死。", loser, winner),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** の筋肉バスターで気絶。", loser, winner),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** の投げたじゃがいもに当たって死亡。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** を魔封波で封印！", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** をちくわソードで成敗🗡", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** をセクシーランジェリーで悩殺💋", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** にウンコを投げつけられて気絶。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** の鼻の穴に致死量の小豆を詰めた🫘", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** のドロップキックで即死。", loser, winner),
	}

	return tmpl[shared.RandInt(1, len(tmpl)+1)-1]
}

// ソロプレイ（無駄アクション）のテンプレートをランダムに取得します。
func GetRandomSoloTmpl() string {
	var tmpl = []string{
		"☀️｜天気が良かったので、 **%s** はお散歩に出かけた。",
		"☀️｜**%s** はナンパにことごとく失敗している。",
		"☀️｜**%s** はキャンプを楽しんでいる。",
		"☀️｜**%s** はバナナの皮をポイ捨てした。誰か引っかかるかな🍌",
		"☀️｜**%s** は豆の収穫をしている。ﾀﾉｼｲ!!",
		"☀️｜**%s** は精神と時の部屋で修行をしている🧘‍♂️",
		"☀️｜**%s** は食べ物を求めて釣りに出かけたが、何も釣れなかった。",
		"☀️｜**%s** は鹿の狩猟に成功しました。",
		"☀️｜**%s** はお尻をポリポリかいている。なんて呑気な。",
		"☀️｜**%s** は立ち止まって花の匂いを嗅いでいる。",
		"☀️｜**%s** は図書館を逆立ちで歩いている。",
	}

	return tmpl[shared.RandInt(1, len(tmpl)+1)-1]
}

// 復活のテンプレートをランダムに取得します
func GetRandomRevivalTmpl(user *discordgo.User) string {
	var tmpl = []string{
		fmt.Sprintf("⚰️｜** %s ** は穢土転生により復活した。", user.Username),
		fmt.Sprintf("⚰️｜** %s ** は往復ビンタで叩き起こされた。復活。", user.Username),
		fmt.Sprintf("⚰️｜** %s ** は神によって蘇生させられた。", user.Username),
	}

	return tmpl[shared.RandInt(1, len(tmpl)+1)-1]
}
