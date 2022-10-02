package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/pkg/router"
)

type (
	unban struct {
	}
)

func (u *unban) bungieHandler(ctx *router.CommandContext) {
	ctx.Reply("Bungie", nil, nil)
}

func (u *unban) discordHandler(ctx *router.CommandContext) {
	ctx.Reply("Discord", nil, nil)
}

func (u *unban) faceitHandler(ctx *router.CommandContext) {
	ctx.Reply("Faceit", nil, nil)
}

func (u *unban) Command() *router.Command {
	cmd := &router.Command{
		Name:        "unban",
		Description: "Unban a user from destinyarena",
	}

	cmd.AddSubCommands(
		&router.SubCommand{
			Name:        "bungie",
			Description: "Unban a user from destinyarena with their bungie id",
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
					Description: "The reason for the unban",
					Required:    true,
				},
			},
			Handler: u.bungieHandler,
		},
		&router.SubCommand{
			Name:        "discord",
			Description: "Unban a user from destinyarena with their discord",
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
					Description: "The reason for the unban",
					Required:    true,
				},
			},
			Handler: u.discordHandler,
		},
		&router.SubCommand{
			Name:        "faceit",
			Description: "Unban a user from destinyarena with their faceit id",
			Options: []*router.CommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "username",
					Description: "The faceit username",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the unban",
					Required:    true,
				},
			},
			Handler: u.faceitHandler,
		},
	)

	return cmd
}
