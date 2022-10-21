package profile

import (
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
	ProfileSummaryID = "profile-summary"
	ProfileTimeoutID = "profile-timeouts"
	ProfileBanID     = "profile-bans"
)

func New() *profile {
	return &profile{}
}

func (p *profile) Components() []router.Component {
	return []router.Component{}
}

func (p *profile) Command() *router.Command {
	return router.NewCommand("profile", "Get a user profile").WithSubCommands(
		router.NewCommand("faceit", "Get a user profile with their faceit name").WithHandler(p.faceitHandler).WithOptions(
			&router.CommandOption{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "username",
				Description: "The faceit username",
				Required:    true,
			},
		),
		router.NewCommand("bungie", "Get a user profile with their bungie name").WithHandler(p.bungieHandler).WithOptions(
			&router.CommandOption{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "The bungie id",
				Required:    true,
			},
		),
		router.NewCommand("discord", "Get a user profile with their discord name").WithHandler(p.discordHandler).WithOptions(
			&router.CommandOption{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "user",
				Description: "The discord user",
				Required:    true,
			},
		),
	)
}

func (p *profile) faceitHandler(ctx *router.CommandContext) {
}

func (p *profile) bungieHandler(ctx *router.CommandContext) {
}

func (p *profile) discordHandler(ctx *router.CommandContext) {
}

//func (p *profile) buildSummary(u *userprofile, ctx *router.Context) (*discordgo.MessageEmbed, []discordgo.MessageComponent) {
//	embed := &discordgo.MessageEmbed{
//		Title: "Profile Summary",
//		Color: SuccessEmbedColor,
//		Thumbnail: &discordgo.MessageEmbedThumbnail{
//			URL: u.DiscordUrl,
//		},
//		Fields: []*discordgo.MessageEmbedField{
//			{
//				Name:   "Discord Username",
//				Value:  fmt.Sprintf("%s#%s", u.DiscordUsername, u.DiscordDiscriminator),
//				Inline: true,
//			},
//			{
//				Name:   "Discord ID",
//				Value:  u.DiscordID,
//				Inline: true,
//			},
//			{
//				Name:   "Bungie Username",
//				Value:  u.BungieUsername,
//				Inline: true,
//			},
//			{
//				Name:   "Bungie ID",
//				Value:  u.BungieID,
//				Inline: true,
//			},
//			{
//				Name:   "Faceit Username",
//				Value:  u.FaceitUsername,
//				Inline: true,
//			},
//			{
//				Name:   "Faceit ID",
//				Value:  u.FaceitID,
//				Inline: true,
//			},
//			{
//				Name:   "Faceit Level",
//				Value:  fmt.Sprintf("%d", u.FaceitLevel),
//				Inline: false,
//			},
//		},
//	}
//
//	if u.Banned {
//		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
//			Name:  "Banned",
//			Value: "Yes",
//		})
//	}
//	//user, _ := ctx.Session.User("411323761116184578")
//	//timeoutButton, err := ctx.Router.GetComponent(ProfileTimeoutID).Build(user)
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//
//	//banButton, err := ctx.Router.GetComponent(ProfileBanID).Build(user)
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	testmenu, err := ctx.Router.GetComponent(TestID).Build()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	profileButtons := []discordgo.MessageComponent{
//		discordgo.ActionsRow{
//			Components: []discordgo.MessageComponent{
//				//timeoutButton,
//				//banButton,
//				testmenu,
//			},
//		},
//	}
//
//	return embed, profileButtons
//}
//
//func (p *profile) timeoutComponentHandler(ctx *router.ComponentContext) {
//	fmt.Println("Timeouts")
//	user := ctx.Args["user"].Value.(*discordgo.User)
//	summaryButton, err := ctx.Router.GetComponent(ProfileSummaryID).Build(user)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	banButton, err := ctx.Router.GetComponent(ProfileBanID).Build(user)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	profileButtons := []discordgo.MessageComponent{
//		discordgo.ActionsRow{
//			Components: []discordgo.MessageComponent{
//				summaryButton,
//				banButton,
//			},
//		},
//	}
//
//	embed := &discordgo.MessageEmbed{
//		Title:       "Profile Timeouts",
//		Color:       SuccessEmbedColor,
//		Description: fmt.Sprintf("Here are the timeouts for %s#%s", user.Username, user.Discriminator),
//		Thumbnail: &discordgo.MessageEmbedThumbnail{
//			URL: user.AvatarURL(""),
//		},
//		Fields: []*discordgo.MessageEmbedField{
//			{
//				Name:  "ID: 454524",
//				Value: "**Time** 4h\n**Reason** Spamming\n**Moderator** <@arturo>\n**Date** 2021-05-01 12:00:00",
//			},
//			{
//				Name:  "ID: 454524",
//				Value: "**Time** 4h\n**Reason** Spamming\n**Moderator** <@arturo>\n**Date** 2021-05-01 12:00:00",
//			},
//			{
//				Name:  "ID: 454524",
//				Value: "**Time** 4h\n**Reason** Spamming\n**Moderator** <@arturo>\n**Date** 2021-05-01 12:00:00",
//			},
//		},
//	}
//
//	err = ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//func (p *profile) banComponentHandler(ctx *router.ComponentContext) {
//	fmt.Println("Bans")
//
//	user := ctx.Args["user"].Value.(*discordgo.User)
//	summaryButton, err := ctx.Router.GetComponent(ProfileSummaryID).Build(user)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	timeoutButtom, _ := ctx.Router.GetComponent(ProfileTimeoutID).Build(user)
//
//	profileButtons := []discordgo.MessageComponent{
//		discordgo.ActionsRow{
//			Components: []discordgo.MessageComponent{
//				summaryButton,
//				timeoutButtom,
//			},
//		},
//	}
//
//	embed := &discordgo.MessageEmbed{
//		Title:       "Profile Bans",
//		Color:       SuccessEmbedColor,
//		Description: fmt.Sprintf("Bans for %s#%s", user.Username, user.Discriminator),
//		Thumbnail: &discordgo.MessageEmbedThumbnail{
//			URL: user.AvatarURL(""),
//		},
//		Fields: []*discordgo.MessageEmbedField{
//			{
//				Name:  "ID: 454524",
//				Value: "**Reason:** Test\n **Moderator:** Test\n **Date:** 2021-09-01 00:00:00",
//			},
//			{
//				Name:  "ID: 454524",
//				Value: "**Reason:** Test\n **Moderator:** Test\n **Date:** 2021-09-01 00:00:00",
//			},
//			{
//				Name:  "ID: 454524",
//				Value: "**Reason:** Test\n **Moderator:** Test\n **Date:** 2021-09-01 00:00:00",
//			},
//		},
//	}
//
//	ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
//}
//func (p *profile) summaryComponentHandler(ctx *router.ComponentContext) {
//	user := ctx.Args["user"].Value.(*discordgo.User)
//
//	u := &userprofile{
//		DiscordID:            user.ID,
//		DiscordUsername:      user.Username,
//		DiscordDiscriminator: user.Discriminator,
//		DiscordUrl:           user.AvatarURL(""),
//		FaceitID:             "123456789",
//		FaceitUsername:       "Destiny",
//		FaceitLevel:          10,
//		BungieUsername:       "Destiny",
//		BungieID:             "123456789",
//		Banned:               false,
//	}
//
//	embed, profileButtons := p.buildSummary(u, &ctx.Context)
//	ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
//
//}
//
//func (p *profile) bungieHandler(ctx *router.CommandContext) {
//	bungieid := ctx.Options["id"].StringValue()
//
//	u := &userprofile{
//		DiscordID:            "123456789",
//		DiscordUsername:      "Destiny",
//		DiscordDiscriminator: "0001",
//		FaceitID:             "123456789",
//		FaceitUsername:       "Destiny",
//		FaceitLevel:          10,
//		BungieUsername:       "Destiny",
//		BungieID:             bungieid,
//		Banned:               false,
//	}
//
//	embed, profileButtons := p.buildSummary(u, ctx.Context)
//
//	ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
//}
//
//func (p *profile) discordHandler(ctx *router.CommandContext) {
//	user, _ := ctx.Session.User(ctx.Options["user"].UserValue(nil).ID)
//
//	u := &userprofile{
//		DiscordID:            user.ID,
//		DiscordUsername:      user.Username,
//		DiscordDiscriminator: user.Discriminator,
//		DiscordUrl:           user.AvatarURL(""),
//		FaceitID:             "123456789",
//		FaceitUsername:       "Destiny",
//		FaceitLevel:          10,
//		BungieUsername:       "Destiny",
//		BungieID:             "123456789",
//		Banned:               false,
//	}
//
//	embed, profileButtons := p.buildSummary(u, ctx.Context)
//	fmt.Println("Discord")
//	err := ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//func (p *profile) faceitHandler(ctx *router.CommandContext) {
//	faceituser := ctx.Options["username"].StringValue()
//
//	u := &userprofile{
//		DiscordID:            "123456789",
//		DiscordUsername:      "Destiny",
//		DiscordDiscriminator: "0001",
//		FaceitID:             "123456789",
//		FaceitUsername:       faceituser,
//		FaceitLevel:          10,
//		BungieUsername:       "Destiny",
//		BungieID:             "123456789",
//		Banned:               false,
//	}
//
//	embed, profileButtons := p.buildSummary(u, ctx.Context)
//	ctx.Reply("", []*discordgo.MessageEmbed{embed}, profileButtons)
//}
//
