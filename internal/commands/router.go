package commands

import "github.com/destinyarena/discord-bot/internal/router"

func New() (router.Router, error) {
	r, _ := router.New()

	r.RegisterCommands(NewBanCommand())
	//r.RegisterCommands(NewProfilesCommand())
	//r.RegisterCommands(NewUnbanCommand())
	//r.RegisterCommands(NewTimeoutCommand())
	//r.RegisterCommands(NewAltsCommand())
	//r.RegisterCommands(NewHistoryCommand())

	return r, nil
}
