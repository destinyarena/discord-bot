package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/destinyarena/discord-bot/pkg/router"
)

// add buttons
type (
	profile struct {
	}

	userprofile struct {
		DiscordID            string
		DiscordUsername      string
		DiscordDiscriminator string
		FaceitID             string
		FaceitUsername       string
		FaceitLevel          int
		BungieID             string
		Banned               bool
	}
)

var (
	profileButtons = []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Timeouts",
					CustomID: "profile-history",
					Style:    discordgo.PrimaryButton,
				},
				discordgo.Button{
					Label:    "Bans",
					CustomID: "profile-bans",
					Style:    discordgo.PrimaryButton,
				},
				discordgo.Button{
					Label:    "Alts",
					CustomID: "profile-alts",
					Style:    discordgo.PrimaryButton,
				},
			},
		},
	}
)

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

func (p *profile) Component() []*router.Component {
	return []*router.Component{
		{
			Name:    "profile-history",
			Type:    discordgo.ButtonComponent,
			Handler: p.historyComponentHandler,
		},
		{
			Name:    "profile-bans",
			Type:    discordgo.ButtonComponent,
			Handler: p.banComponentHandler,
		},
		{
			Name:    "profile-alts",
			Type:    discordgo.ButtonComponent,
			Handler: p.altComponentHandler,
		},
	}
}

func (p *profile) profileReply(ctx *router.Context, u *userprofile) {
	embed := &discordgo.MessageEmbed{
		Title: "Profile",
		Color: SuccessEmbedColor,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Discord Username",
				Value: fmt.Sprintf("%s#%s", u.DiscordUsername, u.DiscordDiscriminator),
			},
			{
				Name:  "Discord ID",
				Value: u.DiscordID,
			},
			{
				Name:  "Faceit Username",
				Value: u.FaceitUsername,
			},
			{
				Name:  "Faceit ID",
				Value: u.FaceitID,
			},
			{
				Name:  "Faceit Level",
				Value: fmt.Sprintf("%d", u.FaceitLevel),
			},
			{
				Name:  "Bungie ID",
				Value: u.BungieID,
			},
		},
	}

	if u.Banned {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Banned",
			Value: "Yes",
		})
	}

	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{embed},
			Components: profileButtons,
		},
	})
}

func (p *profile) historyComponentHandler(ctx *router.ComponentContext) {
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "History",
		},
	})
}

func (p *profile) banComponentHandler(ctx *router.ComponentContext) {
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Ban",
		},
	})
}

func (p *profile) altComponentHandler(ctx *router.ComponentContext) {
	ctx.Session.InteractionRespond(ctx.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Alts",
		},
	})
}

func (p *profile) bungieHandler(ctx *router.Context) {
	bungieid := ctx.Options["id"].StringValue()

	u := &userprofile{
		DiscordID:            "123456789",
		DiscordUsername:      "Destiny",
		DiscordDiscriminator: "0001",
		FaceitID:             "123456789",
		FaceitUsername:       "Destiny",
		FaceitLevel:          10,
		BungieID:             bungieid,
		Banned:               false,
	}

	p.profileReply(ctx, u)
}

func (p *profile) discordHandler(ctx *router.Context) {
	user := ctx.Options["user"].UserValue(nil)

	u := &userprofile{
		DiscordID:            user.ID,
		DiscordUsername:      user.Username,
		DiscordDiscriminator: user.Discriminator,
		FaceitID:             "123456789",
		FaceitUsername:       "Destiny",
		FaceitLevel:          10,
		BungieID:             "123456789",
		Banned:               false,
	}

	p.profileReply(ctx, u)
}

func (p *profile) faceitHandler(ctx *router.Context) {
	faceituser := ctx.Options["username"].StringValue()

	u := &userprofile{
		DiscordID:            "123456789",
		DiscordUsername:      "Destiny",
		DiscordDiscriminator: "0001",
		FaceitID:             "123456789",
		FaceitUsername:       faceituser,
		FaceitLevel:          10,
		BungieID:             "123456789",
		Banned:               false,
	}

	p.profileReply(ctx, u)
}
