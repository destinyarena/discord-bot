package commands

import "github.com/destinyarena/discord-bot/pkg/router"

func New(r router.Router) (router.Router, error) {
	r.RegisterCommands(
		new(ban).Command(),
		new(unban).Command(),
		new(timeout).Command(),
		new(profile).Command(),
	)

	return r, nil
}
