package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/pkg/router"
)

type (
	unban struct {
	}
)

func (u *unban) bungieHandler(ctx *router.Context) {
	ctx.Reply("Bungie")
}

func (u *unban) discordHandler(ctx *router.Context) {
	ctx.Reply("Discord")
}

func (u *unban) faceitHandler(ctx *router.Context) {
	ctx.Reply("Faceit")
}

func (u *unban) Command() *router.Command {
	cmd := &router.Command{
		Name:        "unban",
		Description: "Unban a user from destinyarena",
	}

	cmd.AddSubCommand(
		"bungie",
		"Unban a user from destinyarena with their bungie id",
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
				Description: "The reason for the unban",
				Required:    true,
			},
		},
		u.bungieHandler,
	)

	cmd.AddSubCommand(
		"discord",
		"Unban a user from destinyarena with their discord",
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
				Description: "The reason for the unban",
				Required:    true,
			},
		},
		u.discordHandler,
	)

	cmd.AddSubCommand(
		"faceit",
		"Unban a user from destinyarena with their faceit username",
		[]*router.CommandOption{
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
		u.faceitHandler,
	)

	return cmd
}
