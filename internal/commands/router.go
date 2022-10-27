package commands

import (
	"github.com/destinyarena/discord-bot/internal/commands/profile"
	"github.com/destinyarena/discord-bot/pkg/router"
)

func New(r *router.Router) (*router.Router, error) {
	p := profile.New()
	r.RegisterCommands(
		//new(ban).Command(),
		//new(unban).Command(),
		//new(timeout).Command(),
		p.Command(),
	)

	r.RegisterComponents(
		p.Components()...,
	)

	return r, nil
}
