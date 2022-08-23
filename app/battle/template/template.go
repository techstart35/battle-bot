package template

import (
	"fmt"
	"github.com/techstart35/battle-bot/shared/util"
)

// ソロバトルギミックのテンプレートをランダムに取得します
func GetRandomSoloBattleTmpl(loser string, seed int) string {
	var tmpl = []string{
		fmt.Sprintf("💥｜**%s** は間違えて自爆ボタンを押してしまった💥", loser),
		fmt.Sprintf("💥｜**%s** はバナナの皮で滑って気絶した。", loser),
		fmt.Sprintf("💥｜**%s** は迷子で行方不明になった。", loser),
		fmt.Sprintf("💥｜**%s** は腐った肉を食べて腹痛。戦闘不能。", loser),
		fmt.Sprintf("💥｜**%s** はご飯の食べ過ぎで破裂した。", loser),
		fmt.Sprintf("💥｜**%s** は豆腐の角に頭をぶつけて即死した。", loser),
		fmt.Sprintf("💥｜**%s** はタンスに小指をぶつけてショック死。", loser),
		fmt.Sprintf("💥｜**%s** は鳥のフンが頭に落ちてやる気を失う。脱落。", loser),
		fmt.Sprintf("💥｜**%s** は寝坊で試合に間に合わなかった。", loser),
		fmt.Sprintf("💥｜**%s** は盗んだバイクで走り出したが事故。", loser),
		fmt.Sprintf("💥｜**%s** は夏の暑さで溶けてしまった。", loser),
		fmt.Sprintf("💥｜**%s** は冬の寒さで凍ってしまった。", loser),
		fmt.Sprintf("💥｜**%s** は嫁から鬼電、即帰宅。", loser),
		fmt.Sprintf("💥｜**%s** は飛んでいったカツラを追いかけて退場。", loser),
		fmt.Sprintf("💥｜**%s** は白シャツにカレーを飛ばして戦意喪失。", loser),
		fmt.Sprintf("💥｜**%s** は賞味期限切れの生卵を食べて腹を壊した。", loser),
		fmt.Sprintf("💥｜**%s** は快速特急に乗ってしまい、下車するはずの駅で降りられず。脱落。", loser),
		fmt.Sprintf("💥｜**%s** はランブルの勝ち方を解明すべくアマゾンの奥地へ向かった。", loser),
		fmt.Sprintf("💥｜**%s** は木登りをしていたが、足を滑らせ滑落した。", loser),
		fmt.Sprintf("💥｜**%s** はつま先立ちで歩いていたため、足の指を骨折した。", loser),
		fmt.Sprintf("💥｜**%s** は車に轢かれそうな子供を助けて代わりに事故死。", loser),
		fmt.Sprintf("💥｜**%s** は神隠しにあった。", loser),
		fmt.Sprintf("💥｜**%s** は料理中に指を切って出血死。", loser),
		fmt.Sprintf("💥｜**%s** は転んで頭蓋骨粉砕した。", loser),
		fmt.Sprintf("💥｜**%s** はあつあつのおでんを食べてショック死。", loser),
		fmt.Sprintf("💥｜**%s** はドアに挟まって死んだ。", loser),
		fmt.Sprintf("💥｜**%s** は生け贄に選ばれた。", loser),
	}

	// スライスをシャッフルする
	s := util.ShuffleString(tmpl, seed)

	return s[util.RandInt(1, len(tmpl)+1)-1]
}

// バトルギミックのテンプレートをランダムに取得します
func GetRandomBattleTmpl(winner, loser string, seed int) string {
	var tmpl = []string{
		fmt.Sprintf("⚔️｜👑**%s** は念能力を取得。百式観音を発動し 💀**%s** を駆逐した。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** をブロッコリーで撲殺した🥦", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** に食べられてしまった。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** をタイキックで蹴り倒した。", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** に三輪車で轢かれた。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** に千年殺しをお見舞いした。アーメン。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** を落とし穴に落とした。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** に即効性の毒を盛った。さようなら。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** はラリアットで 💀**%s** を倒した。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は豆鉄砲で 💀**%s** を撃ち抜いた。", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** にきゅうりで殴られ死亡🥒", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は写輪眼を開眼。 💀**%s** は幻術にかけられた。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は北斗百烈拳で 💀**%s** を倒した。ほあたぁ!!", winner, loser),
		fmt.Sprintf("⚔️｜南斗水鳥拳伝承者の 👑**%s** は奥義 飛翔白麗で 💀**%s** を倒した。", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** に秘孔つかれてあべ死。", loser, winner),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** の筋肉バスターで気絶。", loser, winner),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** の投げたじゃがいもに当たって死亡。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** を魔封波で封印した。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** をちくわで撲殺した。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** をセクシーランジェリーで悩殺💋", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** にウ◯コを投げつけられて気絶。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** の鼻の穴に致死量の小豆を詰めた🫘", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** のドロップキックで即死。", loser, winner),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** にパンツを被せられて窒息死。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** を日本刀で一刀両断。", winner, loser),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** にサボテンを投げつけて殺害。", winner, loser),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** にデスノートに名前を書かれて心臓麻痺。", loser, winner),
		fmt.Sprintf("⚔️｜💀**%s** は 👑**%s** の美しさに魅了され、敗けを認めた。", loser, winner),
		fmt.Sprintf("⚔️｜👑**%s** は 💀**%s** を熱湯風呂に落とした。", winner, loser),
	}

	// スライスをシャッフルする
	s := util.ShuffleString(tmpl, seed)

	return s[util.RandInt(1, len(tmpl)+1)-1]
}

// noneのテンプレートをランダムに取得します。
func GetRandomNoneTmpl(winner string, seed int) string {
	var tmpl = []string{
		fmt.Sprintf("☀️｜天気が良かったので、 **%s** はお散歩に出かけた。", winner),
		fmt.Sprintf("☀️｜**%s** はナンパに成功した。", winner),
		fmt.Sprintf("☀️｜**%s** はキャンプを楽しんでいる。", winner),
		fmt.Sprintf("☀️｜**%s** はバナナの皮をポイ捨てした。誰か引っかかるかな🍌", winner),
		fmt.Sprintf("☀️｜**%s** は豆の収穫をしている。ﾀﾉｼｲ!!", winner),
		fmt.Sprintf("☀️｜**%s** は精神と時の部屋で修行をしている🧘‍♂️", winner),
		fmt.Sprintf("☀️｜**%s** は食べ物を求めて釣りに出かけたが、何も釣れなかった。", winner),
		fmt.Sprintf("☀️｜**%s** は鹿の狩猟に成功しました。", winner),
		fmt.Sprintf("☀️｜**%s** はお尻をポリポリかいている。なんて呑気な。", winner),
		fmt.Sprintf("☀️｜**%s** は立ち止まって花の匂いを嗅いでいる。", winner),
		fmt.Sprintf("☀️｜**%s** は図書館を逆立ちで歩いている。", winner),
		fmt.Sprintf("☀️｜**%s** はバイト代、日給1万円を受け取った。", winner),
		fmt.Sprintf("☀️｜**%s** はパズルの最後の1ピースが見つからなくて困っている。", winner),
		fmt.Sprintf("☀️｜**%s** の転職先はOpenSeaに決まった。", winner),
		fmt.Sprintf("☀️｜**%s** は3億円を拾った。...どうする？", winner),
		fmt.Sprintf("☀️｜**%s** は旅行の計画を立てている。どこに行きたい？", winner),
		fmt.Sprintf("☀️｜**%s** は時間を止める能力を手に入れた。", winner),
		fmt.Sprintf("☀️｜**%s** のあだ名は今日から「特攻隊長」だ。", winner),
		fmt.Sprintf("☀️｜**%s** のあだ名は今日から「ババコンガ」だ。", winner),
		fmt.Sprintf("☀️｜**%s** のあだ名は今日から「ダーウィン」だ。", winner),
		fmt.Sprintf("☀️｜**%s** のあだ名は今日から「財閥」だ。", winner),
		fmt.Sprintf("☀️｜**%s** のあだ名は今日から「破壊神」だ。", winner),
		fmt.Sprintf("☀️｜**%s** のあだ名は今日から「マザー」だ。", winner),
		fmt.Sprintf("☀️｜**%s** のあだ名は今日から「ダイナマイト」だ。", winner),
		fmt.Sprintf("☀️｜**%s** は今回のバトルに勝てない。かもしれない。", winner),
		fmt.Sprintf("☀️｜**%s** はゲーセンで大量のメダルをゲットした。", winner),
		fmt.Sprintf("☀️｜**%s** は今時のプリクラが盛れすぎてびっくりしている。", winner),
		fmt.Sprintf("☀️｜**%s** は通勤ラッシュの時間帯にPASMOがはじかれた。", winner),
		fmt.Sprintf("☀️｜**%s** は新刊だと思って漫画を買ったが、すでに持っていた。", winner),
	}

	// スライスをシャッフルする
	s := util.ShuffleString(tmpl, seed)

	return s[util.RandInt(1, len(tmpl)+1)-1]
}

// 復活のテンプレートをランダムに取得します
func GetRandomRevivalTmpl(revival string) string {
	var tmpl = []string{
		fmt.Sprintf("⚰️｜** %s ** は穢土転生により復活した。", revival),
		fmt.Sprintf("⚰️｜** %s ** は往復ビンタで叩き起こされた。復活。", revival),
		fmt.Sprintf("⚰️｜** %s ** は神によって蘇生させられた。", revival),
	}

	return tmpl[util.RandInt(1, len(tmpl)+1)-1]
}
