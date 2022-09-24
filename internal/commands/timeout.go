package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/pkg/router"
)

type (
	timeout struct {
	}
)

func (t *timeout) bungieHandler(ctx *router.Context) {
	ctx.Reply("Bungie")
}

func (t *timeout) discordHandler(ctx *router.Context) {
	ctx.Reply("Discord")
}

func (t *timeout) faceitHandler(ctx *router.Context) {
	ctx.Reply("Faceit")
}

func (t *timeout) Command() *router.Command {
	cmd := &router.Command{
		Name:        "timeout",
		Description: "Timeout a user from destinyarena",
	}

	cmd.AddSubCommand(
		"bungie",
		"Timeout a user from destinyarena with their bungie id",
		[]*router.CommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The bungie id",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The reason for the timeout",
				Required:    true,
			},
		},
		t.bungieHandler,
	)

	cmd.AddSubCommand(
		"discord",
		"Timeout a user from destinyarena with their discord",
		[]*router.CommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The discord user",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The reason for the timeout",
				Required:    true,
			},
		},
		t.discordHandler,
	)

	cmd.AddSubCommand(
		"faceit",
		"Timeout a user from destinyarena with their faceit",
		[]*router.CommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The faceit id",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "reason",
				Description: "The reason for the timeout",
				Required:    true,
			},
		},
		t.faceitHandler,
	)

	return cmd
}
