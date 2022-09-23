package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/internal/router"
)

type (
	ban struct {
		router.BaseCommand
	}

	banFaceit struct {
		router.BaseSubCommandGroup
	}

	banFaceitID struct {
		router.BaseSubCommand
	}

	banFaceitUsername struct {
		router.BaseSubCommand
	}

	banDiscord struct {
		router.BaseSubCommand
	}

	banBungie struct {
		router.BaseSubCommand
	}
)

func (c *ban) Handler(ctx *router.Context) {
	panic("Command not implemented")
}

func (c *banDiscord) Handler(ctx *router.Context) {
	fmt.Println("Reached ban discord sub command")

	user, err := ctx.Session.User(ctx.Options["user"].UserValue(nil).ID)
	if err != nil {
		panic(err)
	}

	reason := ctx.Options["reason"].StringValue()

	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Username: %s Reasion: %s", user.Username, reason),
		},
	})
}

func (c *banFaceit) Handler(ctx *router.Context) {
	fmt.Println("Reached sub ban command")
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "faceit",
		},
	})
}

func (c *banFaceitID) Handler(ctx *router.Context) {
	fmt.Println("Reached ban faceit id sub command")
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "faceit id",
		},
	})
}

func (c *banFaceitUsername) Handler(ctx *router.Context) {
	fmt.Println("Reached ban faceit username sub command")
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "faceit username",
		},
	})
}

func (c *banBungie) Handler(ctx *router.Context) {
	fmt.Println("Reached ban bungie sub command")
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "bungie ban",
		},
	})
}

func NewBanCommand() router.CommandInterface {
	command := &ban{
		router.BaseCommand{
			Name:             "bantest",
			Description:      "Ban a user from destinyarena.",
			SubCommands:      make([]router.SubCommandInterface, 0),
			SubCommandGroups: make([]router.SubCommandGroupInterface, 0),
		},
	}

	command.SubCommandGroups = append(command.SubCommandGroups,
		&banFaceit{
			router.BaseSubCommandGroup{
				Name:        "faceit",
				Description: "Ban a user from faceit.",
				SubCommands: []router.SubCommandInterface{&banFaceitID{
					router.BaseSubCommand{
						Name:        "id",
						Description: "Ban a user by their faceit id.",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "id",
								Description: "Faceit ID",
								Type:        discordgo.ApplicationCommandOptionMentionable,
								Required:    true,
							},
							{
								Name:        "reason",
								Description: "Reason for the ban",
								Type:        discordgo.ApplicationCommandOptionString,
								Required:    true,
							},
						},
					},
				},
					&banFaceitUsername{
						router.BaseSubCommand{
							Name:        "username",
							Description: "Ban a user by their faceit username.",
							Options: []*discordgo.ApplicationCommandOption{
								{
									Name:        "username",
									Description: "Faceit Username",
									Type:        discordgo.ApplicationCommandOptionString,
									Required:    true,
								},
								{
									Name:        "reason",
									Description: "Reason for the ban",
									Type:        discordgo.ApplicationCommandOptionString,
									Required:    true,
								},
							},
						},
					},
				},
			},
		})

	command.SubCommands = append(command.SubCommands,
		&banDiscord{
			router.BaseSubCommand{
				Name:        "discord",
				Description: "Ban a user from destinyarena using discord mention.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "user",
						Description: "User to ban",
						Type:        discordgo.ApplicationCommandOptionUser,
						Required:    true,
					},
					{
						Name:        "reason",
						Description: "Reason for the ban",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
					},
				},
			},
		},
		&banBungie{
			router.BaseSubCommand{
				Name:        "bungie",
				Description: "Ban a user from destinyarena using bungie id.",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "id",
						Description: "Bungie ID",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
					},
					{
						Name:        "reason",
						Description: "Reason for the ban",
						Type:        discordgo.ApplicationCommandOptionString,
						Required:    true,
					},
				},
			},
		})

	return command
}
