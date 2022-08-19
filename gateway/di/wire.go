//go:build wireinject
// +build wireinject

package di

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/wire"
	"github.com/techstart35/battle-bot/app"
	"github.com/techstart35/battle-bot/gateway/persistence/store/battle"
)

// アプリケーションサービスの作成
func InitApp(session *discordgo.Session) (*app.App, error) {
	wire.Build(
		app.NewApp,
		battle.NewRepository,
		wire.Bind(new(app.Repository), new(*battle.Repository)),
	)
	return nil, nil
}

// クエリの作成
func InitQuery() (*battle.Query, error) {
	wire.Build(
		battle.NewQuery,
	)

	return nil, nil
}
