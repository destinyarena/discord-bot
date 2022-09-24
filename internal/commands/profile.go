package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/internal/router"
)

// add buttons
type (
	profile struct {
	}
)

func (p *profile) bungieHandler(ctx *router.Context) {
	ctx.Reply("Bungie")
}

func (p *profile) discordHandler(ctx *router.Context) {
	ctx.Reply("Discord")
}

func (p *profile) faceitHandler(ctx *router.Context) {
	ctx.Reply("Faceit")
}

func (p *profile) Command() *router.Command {
	cmd := &router.Command{
		Name:        "profiles",
		Description: "Get a players profiles",
	}

	cmd.AddSubCommand(
		"faceit",
		"Get a players profile with their faceit username",
		[]*router.CommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "username",
				Description: "The faceit username",
				Required:    true,
			},
		},
		p.faceitHandler,
	)

	cmd.AddSubCommand(
		"discord",
		"Get a players profile with their discord tag",
		[]*router.CommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The discord user",
				Required:    true,
			},
		},
		p.discordHandler,
	)

	cmd.AddSubCommand(
		"bungie",
		"Get a players profile with their bungie id",
		[]*router.CommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The bungie id",
				Required:    true,
			},
		},
		p.bungieHandler,
	)

	return cmd
}
