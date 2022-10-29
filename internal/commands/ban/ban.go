package ban

import (
	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/pkg/router"
)

type (
	ban struct{}
)

func New() *ban {
	return &ban{}
}

func (b *ban) Command() *router.Command {
	return router.NewCommandBuilder("ban", "Ban a user").
		WithSubCommands(
			router.NewCommandBuilder("discord", "Ban a discord user").WithHandler(b.discordHandler).WithOptions(
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The discord user",
					Required:    true,
				},
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the ban",
					Required:    true,
				},
			).MustBuild(),
			router.NewCommandBuilder("faceit", "Ban a faceit user").WithHandler(b.faceitHandler).WithOptions(
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "username",
					Description: "The faceit username",
					Required:    true,
				},
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the ban",
					Required:    true,
				},
			).MustBuild(),
			router.NewCommandBuilder("bungie", "Ban a bungie user").WithHandler(b.bungieHandler).WithOptions(
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "id",
					Description: "The bungie id",
					Required:    true,
				},
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the ban",
					Required:    true,
				},
			).MustBuild(),
		).WithDefaultPermissions(discordgo.PermissionManageServer).
		MustBuild()
}

func (b *ban) discordHandler(ctx *router.CommandContext) {
}

func (b *ban) faceitHandler(ctx *router.CommandContext) {
}

func (b *ban) bungieHandler(ctx *router.CommandContext) {
}
