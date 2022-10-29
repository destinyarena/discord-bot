package unban

import (
	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/pkg/router"
)

type (
	unban struct{}
)

func New() *unban {
	return &unban{}
}

func (u *unban) Command() *router.Command {
	return router.NewCommandBuilder("unban", "Unban a user").
		WithSubCommands(
			router.NewCommandBuilder("discord", "Unban a discord user").WithHandler(u.discordHandler).WithOptions(
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The discord user",
					Required:    true,
				},
			).MustBuild(),
			router.NewCommandBuilder("faceit", "Unban a faceit user").WithHandler(u.faceitHandler).WithOptions(
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "username",
					Description: "The faceit username",
					Required:    true,
				},
			).MustBuild(),
			router.NewCommandBuilder("bungie", "Unban a bungie user").WithHandler(u.bungieHandler).WithOptions(
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "id",
					Description: "The bungie id",
					Required:    true,
				},
			).MustBuild(),
		).
		WithDefaultPermissions(discordgo.PermissionManageServer).
		MustBuild()
}

func (u *unban) discordHandler(ctx *router.CommandContext) {
	// TODO
}

func (u *unban) faceitHandler(ctx *router.CommandContext) {
	// TODO
}

func (u *unban) bungieHandler(ctx *router.CommandContext) {
	// TODO
}
