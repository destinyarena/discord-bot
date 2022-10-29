package commands

import (
	"github.com/arturoguerra/faceitgo"
	"github.com/destinyarena/discord-bot/internal/commands/ban"
	"github.com/destinyarena/discord-bot/internal/commands/profile"
	"github.com/destinyarena/discord-bot/internal/commands/unban"
	"github.com/destinyarena/discord-bot/pkg/router"
)

func New(r *router.Router, faceit *faceitgo.RESTClient) (*router.Router, error) {
	profile := profile.New(faceit)
	unban := unban.New()
	ban := ban.New()

	r.RegisterCommands(
		profile.Command(),
		unban.Command(),
		ban.Command(),
	)

	return r, nil
}
