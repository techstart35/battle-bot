package shared

import "github.com/bwmarrin/discordgo"

// ギルドIDからギルド名を取得します
func GetGuildName(s *discordgo.Session, guildID string) (string, error) {
	guildName := ""
	guild, err := s.Guild(guildID)
	if err != nil {
		return "", CreateErr("ギルドを取得できません", err)
	}
	if guild != nil {
		guildName = guild.Name
	}

	return guildName, nil
}

// ギルドIDとギルド名をマッピングします
var GuildName = map[string]string{
	"984614055681613864":  "TEST SERVER",
	"1008205873610506250": "TEST2",
	"963334640616243201":  "MIRAKO.Community Server",
	"940635506247598180":  "DFT（Dragon Fish Tokyo）",
	"980620726228893756":  "Tokyo Brave Heroes",
	"977497881126789221":  "SleeFi",
	"974182322695991348":  "CryptoNinja Party",
	"961178202871578624":  "Reum House",
	"964047860675010580":  "TUMUGI(KMG進行中)",
	"929660712404549663":  "PSC LAND",
	"942376101215359026":  "Tokyo Alternative Girl",
	"913020213912547339":  "Pixel Heroes DAO",
	"963643228962299984":  "Tokyo NFT LAB",
	"962249385931055134":  "CNF",
	"1008027718438367292": "ユクスエ",
	"894779067981762620":  "1Block Official",
	"901424883202949130":  "CONUSIVERSE",
	"966660223576199179":  "IsekaiBattle",
	"964729260717797438":  "🔻クリプトパス🔻",
	"934283061938503681":  "GEM DEVIL",
	"950526526611419177":  "ペンギン村(元素騎士)",
	"928170795707023380":  "MetaDerby",
	"963384561595711528":  "EGG(ERUMINA GAMING GROUP)",
	"994017569654702100":  "HyggePlus",
	"983938567577423883":  "👻STARING GHOST CREW👻",
	"986276630601281566":  "X SOLDIERS BASE",
	"993168520550551582":  "Gamer Dogs Circle",
}
