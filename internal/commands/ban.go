package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/pkg/router"
)

type (
	ban struct{}
)

func (c *ban) discordHandler(ctx *router.CommandContext) {
	fmt.Println("Reached ban discord sub command")

	user, err := ctx.Session.User(ctx.Options["user"].UserValue(nil).ID)
	if err != nil {
		panic(err)
	}

	reason := ctx.Options["reason"].StringValue()

	ctx.Reply(fmt.Sprintf("Discord ID: %s Discord Username: %s Reasion: %s", user.ID, user.Username, reason), nil, nil)
}

func (c *ban) faceitHandler(ctx *router.CommandContext) {
	fmt.Println("Reached sub ban command")

	user := ctx.Options["user"].StringValue()
	reason := ctx.Options["reason"].StringValue()

	ctx.Reply(fmt.Sprintf("Faceit ID: %s Reasion: %s", user, reason), nil, nil)
}

func (c *ban) bungieHandler(ctx *router.CommandContext) {
	fmt.Println("Reached ban bungie sub command")
	user := ctx.Options["id"].StringValue()
	reason := ctx.Options["reason"].StringValue()
	ctx.Reply(fmt.Sprintf("Bungie ID: %s Reasion: %s", user, reason), nil, nil)
}

func (c *ban) Command() *router.Command {
	command := &router.Command{
		Name:        "ban",
		Description: "Ban a user from the server",
	}

	command.AddSubCommands(
		&router.SubCommand{
			Name:        "discord",
			Description: "Ban a discord user",
			Options: []*router.CommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The user to ban",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the ban",
					Required:    true,
				},
			},
			Handler: c.discordHandler,
		},
		&router.SubCommand{
			Name:        "faceit",
			Description: "Ban a faceit user",
			Options: []*router.CommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "user",
					Description: "The user to ban",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the ban",
					Required:    true,
				},
			},
			Handler: c.faceitHandler,
		},
		&router.SubCommand{
			Name:        "bungie",
			Description: "Ban a bungie user",
			Options: []*router.CommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "id",
					Description: "The user to ban",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "reason",
					Description: "The reason for the ban",
					Required:    true,
				},
			},
			Handler: c.bungieHandler,
		},
	)

	return command
}
