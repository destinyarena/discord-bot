package commands

import "github.com/destinyarena/discord-bot/pkg/router"

func New() (router.Router, error) {
	r, _ := router.New()

	p := new(profile)
	r.RegisterCommands(
		new(ban).Command(),
		p.Command(),
		new(unban).Command(),
		new(timeout).Command(),
	)

	r.RegisterComponents(p.Component()...)

	return r, nil
}
