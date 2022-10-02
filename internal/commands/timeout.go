package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/pkg/router"
)

type (
	timeout struct {
	}
)

func (t *timeout) bungieHandler(ctx *router.CommandContext) {
	ctx.Reply("Bungie", nil, nil)
}

func (t *timeout) discordHandler(ctx *router.CommandContext) {
	ctx.Reply("Discord", nil, nil)
}

func (t *timeout) faceitHandler(ctx *router.CommandContext) {
	ctx.Reply("Faceit", nil, nil)
}

func (t *timeout) Command() *router.Command {
	cmd := &router.Command{
		Name:        "timeout",
		Description: "Timeout a user from destinyarena",
	}

	cmd.AddSubCommands(
		&router.SubCommand{
			Name:        "discord",
			Description: "Timeout a user from destinyarena with their bungie id",
			Options: []*router.CommandOption{
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
			Handler: t.bungieHandler,
		},
		&router.SubCommand{
			Name:        "discord",
			Description: "Timeout a user from destinyarena with their discord",
			Options: []*router.CommandOption{
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
			Handler: t.discordHandler,
		},
		&router.SubCommand{
			Name:        "faceit",
			Description: "Timeout a user from destinyarena with their faceit",
			Options: []*router.CommandOption{
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
			Handler: t.faceitHandler,
		},
	)

	return cmd
}
