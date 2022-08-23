package app

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/battle-bot/domain/model"
	"github.com/techstart35/battle-bot/domain/model/battle"
)

// リポジトリのインターフェイスです
type Repository interface {
	Create(btl *battle.Battle) error
	Update(btl *battle.Battle) error
	Delete(guildID model.GuildID) error
	RejectStart()
}

// クエリーのインターフェイスです
type Query interface {
	FindByGuildID(guildID model.GuildID) (*battle.Battle, error)
	FindAll() ([]*battle.Battle, error)
	IsStartRejected() bool
}

// アプリケーションです
type App struct {
	Repo    Repository
	Query   Query
	Session *discordgo.Session
}

// アプリケーションを作成します
func NewApp(repo Repository, query Query, s *discordgo.Session) *App {
	app := &App{}
	app.Repo = repo
	app.Query = query
	app.Session = s

	return app
}
