package commands

import "github.com/destinyarena/discord-bot/pkg/router"

func New() (router.Router, error) {
	r, _ := router.New()

	r.RegisterCommands(
		new(ban).Command(),
		new(profile).Command(),
		new(unban).Command(),
		new(timeout).Command(),
	)
	//r.RegisterCommands(NewProfilesCommand())
	//r.RegisterCommands(NewUnbanCommand())
	//r.RegisterCommands(NewTimeoutCommand())
	//r.RegisterCommands(NewAltsCommand())
	//r.RegisterCommands(NewHistoryCommand())

	return r, nil
}
