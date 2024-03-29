package commands

import (
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
)

type ban struct {
	*BaseCommand
	gommand.CommandBasics
}

func (c *ban) Init() {
	c.Name = "ban"
	c.Description = "Bans user from discord, faceit and bungie"
	c.PermissionValidators = c.isAllowed(5)
	c.ArgTransformers = []gommand.ArgTransformer{
		{
			Function: gommand.AnyTransformer(gommand.UserTransformer, gommand.StringTransformer),
		},
		{
			Function:  gommand.StringTransformer,
			Remainder: true,
		},
	}
}

func (c *ban) CommandFunction(ctx *gommand.Context) error {
	uid, err := c.GetUID(ctx)
	if err != nil {
		ctx.Reply(err.Error())
		return err
	}

	reason := ctx.Args[1].(string)
	if len(reason) == 0 {
		ctx.Reply("You must provide a ban reason: -ban <fid|did|bid|@|fname> <reason>")
		return nil
	}

	guild := ctx.Session.Guild(disgord.ParseSnowflakeString(c.Config.GuildID))
	if _, err := guild.Get(); err != nil {
		ctx.Reply("Error fetching Guild ID")
		return nil
	}

	// Data fetching
	profile, err := c.Profiles.Get(uid)
	if err != nil {
		ctx.Reply("Error fetching user profile")
		return err
	}

	hubs, err := c.Faceit.GetUserHubs(profile.Faceit)
	if err != nil {
		ctx.Reply("Error fetching user hubs")
		return err
	}

	// Banning
	if err = c.Profiles.Ban(profile.Discord, reason); err != nil {
		ctx.Reply("Error marking user as banned")
		return err
	}

	fields := make([]*disgord.EmbedField, 0)
	errfields := make([]*disgord.EmbedField, 0)

	if derr := guild.Member(disgord.ParseSnowflakeString(profile.Discord)).Ban(&disgord.BanMemberParams{
		Reason: reason,
	}); derr != nil {
		c.Logger.Error(derr)
		errfields = append(errfields, &disgord.EmbedField{
			Name:  "Error",
			Value: "Error banning user from discord server",
		})
	} else {
		fields = append(fields, &disgord.EmbedField{
			Name:  "Successful Discord Ban",
			Value: fmt.Sprintf("<@%s>", profile.Discord),
		})
	}

	for _, hub := range hubs {
		if err = c.Faceit.Ban(hub.HubID, profile.Faceit, reason); err != nil {
			errfields = append(errfields, &disgord.EmbedField{
				Name:  "Error",
				Value: fmt.Sprintf("Error banning user from hub: %s", hub.Name),
			})
		} else {
			fields = append(fields, &disgord.EmbedField{
				Name:  "Successful Hub Ban",
				Value: hub.Name,
			})
		}
	}

	fields = append(fields, errfields...)

	embed := &disgord.Embed{
		Title:       "User Ban",
		Description: fmt.Sprintf("<@%s>", profile.Discord),
		Fields:      fields,
	}

	ctx.Reply(embed)

	return nil
}
