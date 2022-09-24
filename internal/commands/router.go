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

	return r, nil
}
