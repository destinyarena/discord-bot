package commands

import "github.com/destinyarena/discord-bot/internal/router"

type (
	ban struct {
		*router.BasicCommand
	}
)

func (c *ban) Handler(ctx *router.Context) {
	ctx.Reply("This command is not yet implemented.")
}

func NewBanCommand() router.CommandInterface {
	return &ban{
		&router.BasicCommand{
			Name:        "ban",
			Description: "Ban a user from the server.",
		},
	}
}
