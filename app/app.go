package app

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/domain/model/battle"
)

// リポジトリのインターフェイスです
type Repository interface {
	Create(btl battle.Battle) error
	Update(btl battle.Battle) error
	Delete(guildID model.GuildID) error
	RejectStart()
	FindByGuildID(guildID model.GuildID) (battle.Battle, error)
}

// アプリケーションです
type App struct {
	Repo    Repository
	Session *discordgo.Session
}

// アプリケーションを作成します
func NewApp(repo Repository, s *discordgo.Session) *App {
	app := &App{}
	app.Repo = repo
	app.Session = s

	return app
}
