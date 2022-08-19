package guild

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/shared/errors"
)

// ギルドIDからギルド名を取得します
func GetGuildName(s *discordgo.Session, guildID string) (string, error) {
	guildName := ""
	guild, err := s.Guild(guildID)
	if err != nil {
		return "", errors.NewError("ギルドを取得できません", err)
	}
	if guild != nil {
		guildName = guild.Name
	}

	return guildName, nil
}
