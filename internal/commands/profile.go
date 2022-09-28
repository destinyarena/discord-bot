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
	ProfileSummaryID = "profile_summary"
	ProfileTimeoutID = "profile_timeouts"
	ProfileBanID     = "profile_bans"
)

func (p *profile) Command() *router.Command {
	cmd := router.NewCommand(
		"profile",
		"View a users profile",
		nil,
	)

	cmd.AddComponent(&router.Component{
		Name:    ProfileSummaryID,
		Type:    discordgo.ButtonComponent,
		Handler: p.summaryComponentHandler,
		Args: []*router.ComponentArgument{
			{
				Name: "user",
				Type: router.ComponentArgumentTypeUser,
			},
		},
	})

	cmd.AddComponent(&router.Component{
		Name:    ProfileTimeoutID,
		Type:    discordgo.ButtonComponent,
		Handler: p.timeoutComponentHandler,
		Args: []*router.ComponentArgument{
			{
				Name: "user",
				Type: router.ComponentArgumentTypeUser,
			},
		},
	})

	cmd.AddComponent(&router.Component{
		Name:    ProfileBanID,
		Type:    discordgo.ButtonComponent,
		Handler: p.banComponentHandler,
		Args: []*router.ComponentArgument{
			{
				Name: "user",
				Type: router.ComponentArgumentTypeUser,
			},
		},
	})

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

func (p *profile) buildSummary(u *userprofile) (*discordgo.MessageEmbed, []discordgo.MessageComponent) {
	embed := &discordgo.MessageEmbed{
		Title: "Profile Summary",
		Color: SuccessEmbedColor,
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

	profileButtons := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Timeouts",
					CustomID: router.ComponentID(ProfileTimeoutID, u.DiscordID),
					Style:    discordgo.PrimaryButton,
				},
				discordgo.Button{
					Label:    "Bans",
					CustomID: router.ComponentID(ProfileBanID, u.DiscordID),
					Style:    discordgo.PrimaryButton,
				},
			},
		},
	}

	return embed, profileButtons
}

func (p *profile) timeoutComponentHandler(ctx *router.ComponentContext) {
	fmt.Println("Timeouts")
	user := ctx.Args["user"].Value.(*discordgo.User)
	profileButtons := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Summary",
					CustomID: router.ComponentID(ProfileSummaryID, user.ID),
					Style:    discordgo.PrimaryButton,
				},
				discordgo.Button{
					Label:    "Bans",
					CustomID: router.ComponentID(ProfileBanID, user.ID),
					Style:    discordgo.PrimaryButton,
				},
			},
		},
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Profile Timeouts",
		Color:       SuccessEmbedColor,
		Description: fmt.Sprintf("Here are the timeouts for %s#%s", user.Username, user.Discriminator),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: user.AvatarURL(""),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "ID: 454524",
				Value: "**Time** 4h\n**Reason** Spamming\n**Moderator** <@arturo>\n**Date** 2021-05-01 12:00:00",
			},
			{
				Name:  "ID: 454524",
				Value: "**Time** 4h\n**Reason** Spamming\n**Moderator** <@arturo>\n**Date** 2021-05-01 12:00:00",
			},
			{
				Name:  "ID: 454524",
				Value: "**Time** 4h\n**Reason** Spamming\n**Moderator** <@arturo>\n**Date** 2021-05-01 12:00:00",
			},
		},
	}

	err := ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *profile) banComponentHandler(ctx *router.ComponentContext) {
	fmt.Println("Bans")
	user := ctx.Args["user"].Value.(*discordgo.User)
	profileButtons := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "Summary",
					CustomID: router.ComponentID(ProfileSummaryID, user.ID),
					Style:    discordgo.PrimaryButton,
				},
				discordgo.Button{
					Label:    "Timeouts",
					CustomID: router.ComponentID(ProfileTimeoutID, user.ID),
					Style:    discordgo.PrimaryButton,
				},
			},
		},
	}

	embed := &discordgo.MessageEmbed{
		Title:       "Profile Bans",
		Color:       SuccessEmbedColor,
		Description: fmt.Sprintf("Bans for %s#%s", user.Username, user.Discriminator),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: user.AvatarURL(""),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "ID: 454524",
				Value: "**Reason:** Test\n **Moderator:** Test\n **Date:** 2021-09-01 00:00:00",
			},
			{
				Name:  "ID: 454524",
				Value: "**Reason:** Test\n **Moderator:** Test\n **Date:** 2021-09-01 00:00:00",
			},
			{
				Name:  "ID: 454524",
				Value: "**Reason:** Test\n **Moderator:** Test\n **Date:** 2021-09-01 00:00:00",
			},
		},
	}

	ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
}
func (p *profile) summaryComponentHandler(ctx *router.ComponentContext) {
	user := ctx.Args["user"].Value.(*discordgo.User)

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

	embed, profileButtons := p.buildSummary(u)
	ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)

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
		BungieUsername:       "Destiny",
		BungieID:             bungieid,
		Banned:               false,
	}

	embed, profileButtons := p.buildSummary(u)
	ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
}

func (p *profile) discordHandler(ctx *router.Context) {
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

	embed, profileButtons := p.buildSummary(u)
	ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
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
		BungieUsername:       "Destiny",
		BungieID:             "123456789",
		Banned:               false,
	}

	embed, profileButtons := p.buildSummary(u)
	ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
}
