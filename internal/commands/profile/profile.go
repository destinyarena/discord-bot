package profile

import (
	"fmt"

	"github.com/arturoguerra/faceitgo"
	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/pkg/router"
)

// add buttons
type (
	profile struct {
		Faceit *faceitgo.RESTClient
	}

	userprofile struct {
		DiscordID            string
		DiscordUsername      string
		DiscordDiscriminator string
		DiscordUrl           string
		FaceitID             string
		FaceitUsername       string
		FaceitLevel          int
		BungieUsername       string
		BungieID             string
		Banned               bool
	}
)

const (
	HistoryButtonID     = "/profile/history/%s/%s"
	HistoryButtonPageID = "/profile/history/%s/%s/%s"
	SuccessEmbedColor   = 0xF30707
)

func New(f *faceitgo.RESTClient) *profile {
	return &profile{f}
}

func (p *profile) Command() *router.Command {
	return router.NewCommandBuilder("profile", "Get a user profile").
		WithSubCommands(
			router.NewCommandBuilder("faceit", "Get a user profile").WithHandler(p.faceitHandler).WithOptions(
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "username",
					Description: "The faceit username",
					Required:    true,
				},
			).MustBuild(),
			router.NewCommandBuilder("bungie", "Get a user profile").WithHandler(p.bungieHandler).WithOptions(
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "id",
					Description: "The bungie id",
					Required:    true,
				},
			).MustBuild(),
			router.NewCommandBuilder("discord", "Get a user profile").WithHandler(p.discordHandler).WithOptions(
				&router.CommandOption{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user",
					Description: "The discord user",
					Required:    true,
				},
			).MustBuild(),
		).
		WithDefaultPermissions(discordgo.PermissionManageServer).
		MustBuild()
}

func (p *profile) Components() []*router.Component {
	return []*router.Component{}
}

func (p *profile) bungieHandler(ctx *router.CommandContext) {
	bungieid := ctx.Options["id"].StringValue()

	u := &userprofile{
		DiscordID:            "123456789",
		DiscordUsername:      "Destiny",
		DiscordDiscriminator: "0001",
		FaceitID:             "123456789",
		FaceitUsername:       "Destiny",
		FaceitLevel:          10,
		BungieUsername:       "Destiny",
		BungieID:             bungieid,
		Banned:               false,
	}

	embed, profileButtons := p.buildSummary(u, ctx.Context)

	ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
}

func (p *profile) discordHandler(ctx *router.CommandContext) {
	user, _ := ctx.Session.User(ctx.Options["user"].UserValue(nil).ID)

	u := &userprofile{
		DiscordID:            user.ID,
		DiscordUsername:      user.Username,
		DiscordDiscriminator: user.Discriminator,
		DiscordUrl:           user.AvatarURL(""),
		FaceitID:             "123456789",
		FaceitUsername:       "Destiny",
		FaceitLevel:          10,
		BungieUsername:       "Destiny",
		BungieID:             "123456789",
		Banned:               false,
	}

	embed, profileButtons := p.buildSummary(u, ctx.Context)
	fmt.Println("Discord")
	err := ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *profile) faceitHandler(ctx *router.CommandContext) {
	faceituser := ctx.Options["username"].StringValue()
	fmt.Println("Getting user ", faceituser)

	user, err := p.Faceit.GetPlayer(faceituser, "", "")
	if err != nil {
		fmt.Println("Error getting player", err)
		return
	}

	fmt.Println("Got user ", user)

	level := 0

	if game, ok := user.Games["destiny2"]; ok {
		level = game.SkillLevel
	}

	u := &userprofile{
		DiscordID:            "123456789",
		DiscordUsername:      "Destiny",
		DiscordDiscriminator: "0001",
		FaceitID:             user.PlayerID,
		FaceitUsername:       user.Nickname,
		FaceitLevel:          level,
		BungieUsername:       "Destiny",
		BungieID:             "123456789",
		Banned:               false,
	}

	embed, buttons := p.buildSummary(u, ctx.Context)
	ctx.Reply("", []*discordgo.MessageEmbed{embed}, buttons)
}

func (p *profile) buildSummary(u *userprofile, ctx *router.Context) (*discordgo.MessageEmbed, []discordgo.MessageComponent) {
	embed := &discordgo.MessageEmbed{
		Title: "Profile Summary",
		Color: int(discordgo.SuccessButton),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: u.DiscordUrl,
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Discord Username",
				Value:  fmt.Sprintf("%s#%s", u.DiscordUsername, u.DiscordDiscriminator),
				Inline: true,
			},
			{
				Name:   "Discord ID",
				Value:  u.DiscordID,
				Inline: true,
			},
			{
				Name:   "Bungie Username",
				Value:  u.BungieUsername,
				Inline: true,
			},
			{
				Name:   "Bungie ID",
				Value:  u.BungieID,
				Inline: true,
			},
			{
				Name:   "Faceit Username",
				Value:  u.FaceitUsername,
				Inline: true,
			},
			{
				Name:   "Faceit ID",
				Value:  u.FaceitID,
				Inline: true,
			},
			{
				Name:   "Faceit Level",
				Value:  fmt.Sprintf("%d", u.FaceitLevel),
				Inline: false,
			},
		},
	}

	if u.Banned {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Banned",
			Value: "Yes",
		})
	}

	return embed, nil
}
