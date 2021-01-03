package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	faceit "github.com/destinyarena/bot/pkg/faceit"
	profiles "github.com/destinyarena/bot/pkg/profiles"
	"github.com/destinyarena/bot/pkg/router"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

func (c *Commands) ban(ctx *router.Context) {
	// Setup logic
	reason := strings.Join(ctx.Args[1:], " ")
	if len(reason) == 0 {
		ctx.Reply("You must provide a ban reason: -ban <fid|did|bid|@|fname> <reason>")
		return
	}

	guild, err := ctx.Session.Guild(c.Config.Discord.GuildID)
	if err != nil {
		ctx.Reply("Error fetching Guild ID")
		return
	}

	uid, err := c.GetUID(ctx)
	if err != nil {
		ctx.Reply(err.Error())
		return
	}

	// Connections
	fconn, err := grpc.Dial(c.Config.GRPC.Faceit, grpc.WithInsecure())
	if err != nil {
		ctx.Reply("Error connecting to internal faceit system")
		return
	}

	defer fconn.Close()

	fClient := faceit.NewFaceitClient(fconn)

	pconn, err := grpc.Dial(c.Config.GRPC.Profile, grpc.WithInsecure())
	if err != nil {
		ctx.Reply("Error conneting to internal profiles system")
		return
	}
	defer pconn.Close()

	pClient := profiles.NewProfilesClient(pconn)

	// Data fetching
	profile, err := c.getProfile(uid)
	if err != nil {
		ctx.Reply("Error fetching user profile")
		return
	}

	_, err = ctx.Session.User(profile.Discord)
	if err != nil {
		ctx.Reply("Error fetching discord user")
		return
	}

	hubs, err := c.GetHubs(profile.Faceit)
	if err != nil {
		ctx.Reply("Error fetching user hubs")
		return
	}

	errfields := make([]*discordgo.MessageEmbedField, 0)
	fields := make([]*discordgo.MessageEmbedField, 0)

	// Banning
	_, err = pClient.Ban(context.Background(), &profiles.BanRequest{
		Id:     profile.Discord,
		Reason: reason,
	})
	if err != nil {
		ctx.Reply("Error marking user as banned")
		return
	}

	if derr := ctx.Session.GuildBanCreateWithReason(guild.ID, profile.Discord, reason, 7); derr != nil {
		log.Error(derr)
		errfields = append(errfields, &discordgo.MessageEmbedField{
			Name:  "Error",
			Value: "Error banning user from discord server",
		})
	} else {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  "Successful Discord Ban",
			Value: fmt.Sprintf("<@%s>", profile.Discord),
		})
	}

	for _, hub := range hubs {
		if _, err = fClient.Ban(context.Background(), &faceit.BanRequest{
			Hubid:  hub.Hubid,
			Guid:   profile.Faceit,
			Reason: reason,
		}); err != nil {
			errfields = append(errfields, &discordgo.MessageEmbedField{
				Name:  "Error",
				Value: fmt.Sprintf("Error banning user from hub: %s", hub.Name),
			})
		} else {
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  "Successful Hub Ban",
				Value: fmt.Sprintf("%s", hub.Name),
			})
		}
	}

	fields = append(fields, errfields...)

	embed := &discordgo.MessageEmbed{
		Title:       "User Ban",
		Description: fmt.Sprintf("<@%s>", profile.Discord),
		Fields:      fields,
	}

	ctx.ReplyEmbed(embed)
	return
}
