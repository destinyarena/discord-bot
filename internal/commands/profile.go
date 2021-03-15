package commands

import (
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/auttaja/gommand"
	"github.com/destinyarena/bot/internal/faceit"
	"github.com/labstack/gommon/log"
)

type profile struct {
	*BaseCommand
	gommand.CommandBasics
}

func (c *profile) Init() {
	c.Name = "profile"
	c.Description = "Gets a members profile"
	//c.PermissionValidators = []func(*gommand.Context) (string, bool){
	//	c.isOwner(),
	//	c.isAdmin(),
	//	c.isMod(),
	//}
}

func getBannedValue(banned bool) string {
	if banned {
		return "yes"
	}

	return "no"
}

func (c *profile) CommandFunction(ctx *gommand.Context) error {
	var embed *disgord.Embed
	var uid string

	uid, err := c.GetUID(ctx)
	if err != nil {
		embed = &disgord.Embed{
			Title:       "Error fetching Profile",
			Description: err.Error(),
			Color:       0xF30707,
		}

		ctx.Reply(embed)
		return err
	}

	c.Logger.Infof("UID: %s", uid)

	profile, err := c.Profiles.Get(uid)
	if err != nil {
		embed = &disgord.Embed{
			Title:       "Error Fetching User",
			Description: "Looks like you provided an invaild id or the user never registered",
			Color:       0xF30707,
		}

		ctx.Reply(embed)
		return err
	}

	guild := ctx.Session.Guild(disgord.ParseSnowflakeString(c.Config.GuildID))
	duser, err := guild.Member(disgord.ParseSnowflakeString(profile.Discord)).Get()
	if err != nil {
		c.Logger.Error(err)
		embed = &disgord.Embed{
			Title:       "Error Fetching User",
			Description: "Error fetching Profile",
			Color:       0xf30707,
		}
		ctx.Reply(embed)
		return err
	}
	c.Logger.Info("PAIN")

	fprofile, err := c.Faceit.GetProfileByID(profile.Faceit)
	if err != nil {
		c.Logger.Error(err)
		fprofile = &faceit.Profile{
			GUID:     profile.Faceit,
			Username: "Unavailable",
			Level:    3,
		}
	}
	c.Logger.Info("PAIN")

	fields := make([]*disgord.EmbedField, 0)

	fields = append(fields, &disgord.EmbedField{
		Name:  "Discord Username",
		Value: fmt.Sprintf("%s#%s", duser.User.Username, duser.User.Discriminator),
	})

	fields = append(fields, &disgord.EmbedField{
		Name:  "Discord ID",
		Value: duser.UserID.String(),
	})

	fields = append(fields, &disgord.EmbedField{
		Name:  "Bungie ID",
		Value: profile.Bungie,
	})

	fields = append(fields, &disgord.EmbedField{
		Name:  "Faceit Username",
		Value: fprofile.Username,
	})

	fields = append(fields, &disgord.EmbedField{
		Name:  "Faceit GUID",
		Value: fprofile.GUID,
	})

	fields = append(fields, &disgord.EmbedField{
		Name:  "Faceit Skill Level",
		Value: fmt.Sprintf("%d", fprofile.Level),
	})

	fields = append(fields, &disgord.EmbedField{
		Name:  "Banned",
		Value: getBannedValue(profile.Banned),
	})

	if profile.Banned && len(profile.BanReason) > 0 {
		fields = append(fields, &disgord.EmbedField{
			Name:  "Ban Reason",
			Value: profile.BanReason,
		})
	}

	log.Info("Trying to fetch user hubs")
	hubs, err := c.Faceit.GetUserHubs(profile.Faceit)
	if err != nil {
		log.Error(err)
	} else {
		for _, hub := range hubs {
			fields = append(fields, &disgord.EmbedField{
				Name:  fmt.Sprintf("Hub Name: %s", hub.Name),
				Value: fmt.Sprintf("Hub ID: %s \n Game ID: %s", hub.HubID, hub.GameID),
			})
		}
	}

	embed = &disgord.Embed{
		Title:  "User Profile",
		Color:  0x019fd8,
		Fields: fields,
	}

	ctx.Reply(embed)
	return nil
}
